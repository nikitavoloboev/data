# Neural networks
[This](https://www.youtube.com/watch?v=aircAruvnKk&feature=youtu.be) is a great introduction. Still learning.

## Notes
- Neural Networks are really good at identifying patterns in data. As a classic example, if you wanted to predict housing prices, you could build a data set that maps features about houses (square feet, location, proximity to Caltrain, etc) onto their actual price, and then train a network to recognize the complex relationship between features and pricing. Training happens by feeding the network features, letting it make a guess about the price, and then correcting the guess (backpropagation).
	- Convolutional Neural Networks work similarly, but with images. Instead of giving a CNN discrete features, you'll usually just use the pixels of the image itself. Through a series of layers, the CNN is able to build features itself (traditionally things like edges, corners) and learn patterns in image data. For example, a CNN might be trained on a dataset that maps images onto labels, and learn how to label new images on its own.

## Links
- [Capsule Networks (CapsNets) – Tutorial](https://www.youtube.com/watch?v=pPN8d0E3900)
- [Chris Olah explains neural nets](https://www.youtube.com/watch?v=vdqu6fvjc5c)