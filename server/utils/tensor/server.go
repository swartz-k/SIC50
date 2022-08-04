package tensor

import (
	"fmt"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/pkg/errors"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"gocv.io/x/gocv"
	"image"
	"path"
)

var model *tf.SavedModel

//
func getModel() (*tf.SavedModel, error) {
	var err error
	mPath := path.Join(config.Cfg.WorkDir, "model_p")
	if model != nil {
		return model, nil
	}
	model, err = tf.LoadSavedModel(mPath, []string{"serve"}, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "load model from %s", mPath)
	}
	return model, nil
}

const (
	H, W = 201, 201
)

func Call(imgPath string) ([]*tf.Tensor, error) {
	tensor, err := makeTensorFromImage(imgPath)
	//tensor, err := makeImageTensorByGocv(imgPath)
	if err != nil {
		return nil, errors.Wrap(err, "new tensor failed")
	}

	input, output := getInputOutput("serving_default_input_input", "StatefulPartitionedCall")
	model, err := getModel()
	if err != nil {
		return nil, err
	}
	r, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			input.Output(0): tensor,
		},
		[]tf.Output{
			output.Output(0),
		},
		nil)
	return r, err
}

func constructGraphToNormalizeImage() (graph *tf.Graph, input, output_img tf.Output, err error) {
	// - 这个模型需要输入图片的格式为 198(h)*198(w)
	// - 图像的通道信息应该为3通道
	// - 需要将图像进行灰度化处理
	const (
		Mean  = float32(117)
		Scale = float32(3)
	)
	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	// - 这里其实就是对图片像素进行了处理

	// - 利用双线性插值的方法改变图片的大小
	output_img = op.ResizeBilinear(s,
		// - 将tensor扩展为4维，加入了batch_size这个维度
		op.ExpandDims(s,
			//  - 强制转化类型,注意设置Channels为1，即做了图像的灰度化处理
			op.Cast(s, op.DecodePng(s, input, op.DecodePngChannels(3)), tf.Float),
			op.Const(s.SubScope("make_batch"), int32(0))),
		op.Const(s.SubScope("size"), []int32{H, W}))

	graph, err = s.Finalize()
	return graph, input, output_img, err
}

/*
   该函数传入图像路径，读入图片，并将图像转化为tensor
   并将图像处理成model需要的输入格式
   返回的Tensor格式为：[batch_size,height，width，channels]
*/
func makeTensorFromImage(filename string) (*tf.Tensor, error) {
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return nil, errors.New("can't read image!")
	}
	// 转化为需要的图像大小
	gocv.Resize(img, &img, image.Pt(H, W), 0, 0, gocv.InterpolationNearestNeighbor)

	img_tmp := gocv.NewMat()
	img.ConvertTo(&img_tmp, gocv.MatTypeCV32FC3)
	defer img_tmp.Close()
	log.Info("img tmp size %+v", img_tmp.Size())
	//img_slider是图像Mat转化成的切片
	img_slider, err := img_tmp.DataPtrFloat32() //Mat转slice切片
	if err != nil {
		log.Info("slider %+v", err)
		return nil, err
	}
	// slice切片转成tensor
	tensor, err := tf.NewTensor(img_slider)
	if err != nil {
		log.Info("err %+v", err)
		return nil, err
	}

	// Construct a graph to normalize the image
	graph, input, output, err := constructGraphTo4DImage()
	if err != nil {
		fmt.Println("constructGraphToNormalizeImage is failed!")
		return nil, err
	}
	// Execute that graph to normalize this one image
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		fmt.Println("New Session is failed!")
		return nil, err
	}
	defer session.Close()
	// 开启session得到处理后的图像tensor
	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		fmt.Println("Run Session is failed!")
		fmt.Println(err)
		return nil, err
	}
	return normalized[0], nil
}

/*
   定义操作图
   将一维的tensor变量reshape成四维的tensor，可用于模型的输入
*/
func constructGraphTo4DImage() (graph *tf.Graph, input, out_put_4d tf.Output, err error) {
	s := op.NewScope()
	input = op.Placeholder(s, tf.Float)
	out_put_4d = op.Reshape(s, input,
		op.Const(s.SubScope("shape"), []int32{-1, H, W, 3}))
	graph, err = s.Finalize()
	return graph, input, out_put_4d, err
}

func getInputOutput(input, output string) (*tf.Operation, *tf.Operation) {

	i := model.Graph.Operation(input)
	if i != nil {
		fmt.Printf("iiii %s input %d, output %d \n", i.Name(), i.NumInputs(), i.NumOutputs())
	} else {
		fmt.Printf("input op %s nil\n", input)
	}
	o := model.Graph.Operation(output)
	if o != nil {
		fmt.Printf("oooo %s input %d, output %d \n", o.Name(), o.NumInputs(), o.NumOutputs())
	} else {
		fmt.Printf("output op %s nil\n", output)
	}
	return i, o
}

func makeImageTensorByGocv(filename string) (*tf.Tensor, error) {
	const (
		CHANEELS = 1
	)
	// 读取图像
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return nil, errors.New("can't read image!")
	}

	// 转化为灰度图像
	//gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
	// 转化为需要的图像大小
	gocv.Resize(img, &img, image.Pt(H, W), 0, 0, gocv.InterpolationCubic)

	img_tmp := gocv.NewMat()
	img.ConvertTo(&img_tmp, gocv.MatTypeCV32FC3)
	//img.ConvertTo(&img_tmp, gocv.MatTypeCV32FC1)
	defer img_tmp.Close()
	//log.Info("img tmp size %+v", img_tmp.Size())
	//img_slider是图像Mat转化成的切片
	img_slider, err := img_tmp.DataPtrFloat32() //Mat转slice切片
	if err != nil {
		log.Info("slider %+v", err)
		return nil, err
	}
	// slice切片转成tensor
	tensor, err := tf.NewTensor(img_slider)
	if err != nil {
		log.Info("err %+v", err)
		return nil, err
	}

	graph, input, output, err := constructGraphTo4DImage()
	if err != nil {
		fmt.Println("constructGraph4DImage is failed!")
		return nil, err
	}
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		fmt.Println("New Session is failed!")
		return nil, err
	}
	defer session.Close()
	//Run得到四维的tensor
	img_4d, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		fmt.Println("Run Session is failed!")
		fmt.Println(err)
		return nil, err
	}
	println(img_4d[0].Shape())
	return img_4d[0], nil
}

func Cal(imgPath, input, output string) ([]*tf.Tensor, error) {
	//for _, i := range model.Graph.Operations() {
	//	fmt.Printf("operation name: %s, inputs: %d, outputs: %d, type: %s \n", i.Name(), i.NumInputs(), i.NumOutputs(), i.Type())
	//}
	var tensor *tf.Tensor
	//err := image.Transfer(imgPath)
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
	//img_tmp := gocv.NewMat()
	//img.ConvertTo(&img_tmp, gocv.MatTypeCV32FC1)
	//defer img_tmp.Close()
	//img_slider是图像Mat转化成的切片
	//gocv.Resize(img_tmp, &img_tmp, image.Point{X: 0, Y: 0}, 198, 198, gocv.InterpolationNearestNeighbor)
	//
	//log.Info("reshape size %+v", img_tmp)
	//content := img_tmp.ToBytes()

	resizeImage := gocv.NewMat()
	gocv.Resize(img, &resizeImage, image.Point{X: H, Y: W}, 0, 0, gocv.InterpolationNearestNeighbor)
	tensor, err := tf.NewTensor(resizeImage.ToBytes())
	if err != nil {
		return nil, errors.Wrap(err, "deal with image tra")
	}

	//
	//tensor, err := tf.NewTensor([1][198][198][3]float32{})
	//if err != nil {
	//	return nil, err
	//}

	//log.Info("deal with image %s", imgPath)
	i := model.Graph.Operation(input)
	if i != nil {
		fmt.Printf("iiii %s input %d, output %d \n", i.Name(), i.NumInputs(), i.NumOutputs())
	} else {
		fmt.Printf("input op %s nil\n", input)
	}
	o := model.Graph.Operation(output)
	if o != nil {
		fmt.Printf("oooo %s input %d, output %d \n", o.Name(), o.NumInputs(), o.NumOutputs())
	} else {
		fmt.Printf("output op %s nil\n", output)
	}

	if model.Session == nil {
		return nil, fmt.Errorf("model session is nil")
	}
	if tensor == nil {
		return nil, fmt.Errorf("tensor is nil")
	}
	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			//model.Graph.Operation("conv2d/kernel").Output(0): tensor, // Replace this with your input layer name
			i.Output(0): tensor, // Replace this with your input layer name
		},
		[]tf.Output{
			o.Output(0), // Replace this with your output layer name
		},
		nil,
	)
	log.Info("cal re %+v, err %+v", result, err)
	return result, err
}
