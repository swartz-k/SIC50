import tensorflow as tf
import keras_preprocessing
from keras_preprocessing import image
from keras_preprocessing.image import ImageDataGenerator
import sys

TRAINING_DIR = sys.argv[1]
epoch = sys.argv[2]
steps_per_epoch = sys.argv[3]

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


VALIDATION_DIR = sys.argv[1]
# VALIDATION_DIR = ('../images/Cephalotaxin')
validation_datagen = ImageDataGenerator(rescale = 1./255)

train_generator = training_datagen.flow_from_directory(
	TRAINING_DIR,
	target_size=(201,201),
	class_mode='binary',
  batch_size=2
)

validation_generator = validation_datagen.flow_from_directory(
	VALIDATION_DIR,
	target_size=(201,201),
	class_mode='binary',
  batch_size=2
)

model = tf.keras.models.Sequential([
    # This is the first convolution
    tf.keras.layers.Conv2D(64, (3,3), activation='relu', input_shape=(201, 201, 3)),
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
    tf.keras.layers.Dense(1, activation='sigmoid')
])


model.summary()

model.compile(loss = tf.keras.losses.BinaryCrossentropy(),
              optimizer = tf.keras.optimizers.Adam(learning_rate=1e-3),
              metrics = [tf.keras.metrics.BinaryAccuracy(),
                         tf.keras.metrics.FalseNegatives()])

# fixme: dynamic images = epochs * steps_per_epoch
history = model.fit(train_generator, epochs=int(epoch), steps_per_epoch=int(steps_per_epoch),
                    validation_data = validation_generator, verbose = 1, validation_steps=int(epoch))

print(history)


print(model.evaluate(validation_generator))

