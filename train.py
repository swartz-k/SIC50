import tensorflow as tf
import keras_preprocessing
from keras_preprocessing import image
from keras_preprocessing.image import ImageDataGenerator

"""
Model: "sequential"
_________________________________________________________________
Layer (type)                 Output Shape              Param #
=================================================================
input (Conv2D)               (None, 196, 196, 64)      1792
_________________________________________________________________
max_pooling2d (MaxPooling2D) (None, 98, 98, 64)        0
_________________________________________________________________
conv2d (Conv2D)              (None, 96, 96, 64)        36928
_________________________________________________________________
max_pooling2d_1 (MaxPooling2 (None, 48, 48, 64)        0
_________________________________________________________________
conv2d_1 (Conv2D)            (None, 46, 46, 128)       73856
_________________________________________________________________
max_pooling2d_2 (MaxPooling2 (None, 23, 23, 128)       0
_________________________________________________________________
conv2d_2 (Conv2D)            (None, 21, 21, 128)       147584
_________________________________________________________________
max_pooling2d_3 (MaxPooling2 (None, 10, 10, 128)       0
_________________________________________________________________
flatten (Flatten)            (None, 12800)             0
_________________________________________________________________
dropout (Dropout)            (None, 12800)             0
_________________________________________________________________
dense (Dense)                (None, 512)               6554112
_________________________________________________________________
output (Dense)               (None, 2)                 1026
=================================================================
Total params: 6,815,298
Trainable params: 6,815,298
Non-trainable params: 0

"""
TRAINING_DIR = ('../images/success')
# TRAINING_DIR = ('../images/Cephalotaxin')
training_datagen = ImageDataGenerator(
      rescale = 1./255,
	    rotation_range=40,
      width_shift_range=0.2,
      height_shift_range=0.2,
      shear_range=0.2,
      zoom_range=0.2,
      horizontal_flip=True,
      fill_mode='nearest')


VALIDATION_DIR = ('../images/success')
# VALIDATION_DIR = ('../images/Cephalotaxin')
validation_datagen = ImageDataGenerator(rescale = 1./255)

train_generator = training_datagen.flow_from_directory(
	TRAINING_DIR,
	target_size=(198,198),
	class_mode='categorical',
  batch_size=2
)

validation_generator = validation_datagen.flow_from_directory(
	VALIDATION_DIR,
	target_size=(198,198),
	class_mode='categorical',
  batch_size=2
)

model = tf.keras.models.Sequential([
    # Note the input shape is the desired size of the image 150x150 with 3 bytes color
    # This is the first convolution
    tf.keras.layers.Conv2D(64, (3,3), activation='relu', input_shape=(198, 198, 3), name='input'),
    tf.keras.layers.MaxPooling2D(2, 2),
    # The second convolution
    tf.keras.layers.Conv2D(64, (3,3), activation='relu'),
    tf.keras.layers.MaxPooling2D(2,2),
    # The third convolution
    tf.keras.layers.Conv2D(128, (3,3), activation='relu'),
    tf.keras.layers.MaxPooling2D(2,2),
    # The fourth convolution
    tf.keras.layers.Conv2D(128, (3,3), activation='relu'),
    tf.keras.layers.MaxPooling2D(2,2),
    # Flatten the results to feed into a DNN
    tf.keras.layers.Flatten(),
    tf.keras.layers.Dropout(0.5),
    # 512 neuron hidden layer
    tf.keras.layers.Dense(512, activation='relu'),
    tf.keras.layers.Dense(2, activation='softmax', name='output')
])


model.summary()

model.compile(loss = 'categorical_crossentropy', optimizer='rmsprop', metrics=['accuracy'])

history = model.fit(train_generator, epochs=30, steps_per_epoch=30, validation_data = validation_generator, verbose = 1, validation_steps=3)

print(history)

model.save('model_p')