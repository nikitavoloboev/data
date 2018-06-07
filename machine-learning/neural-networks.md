# Neural networks
[This](https://www.youtube.com/watch?v=aircAruvnKk&feature=youtu.be) is a great introduction.

## Notes
- Neural Networks are great identifying patterns in data. As a classic example, if you wanted to predict housing prices, you could build a data set that maps features about houses (square feet, location, proximity to Caltrain, etc) onto their actual price, and then train a network to recognize the complex relationship between features and pricing. Training happens by feeding the network features, letting it make a guess about the price, and then correcting the guess (backpropagation).
	- Convolutional Neural Networks work similarly, but with images. Instead of giving a CNN discrete features, you'll usually just use the pixels of the image itself. Through a series of layers, the CNN is able to build features itself (traditionally things like edges, corners) and learn patterns in image data. For example, a CNN might be trained on a dataset that maps images onto labels, and learn how to label new images on its own.

## Links
- [A Neural Network Playground](https://playground.tensorflow.org)
- [A Beginner's Guide To Understanding Convolutional Neural Networks](https://adeshpande3.github.io/adeshpande3.github.io/A-Beginner's-Guide-To-Understanding-Convolutional-Neural-Networks/)
- [Capsule Networks (CapsNets) – Tutorial](https://www.youtube.com/watch?v=pPN8d0E3900)
- [Chris Olah explains neural nets](https://www.youtube.com/watch?v=vdqu6fvjc5c)
- [How I Shipped a Neural Network on iOS with CoreML, PyTorch, and React Native](https://attardi.org/pytorch-and-coreml) - Detailed and awesome article.
- [Generative Adversarial Networks (GANs) in 50 lines of code (PyTorch)](https://medium.com/@devnag/generative-adversarial-networks-gans-in-50-lines-of-code-pytorch-e81b79659e3f)
- [Neural Networks, Types, and Functional Programming](http://colah.github.io/posts/2015-09-NN-Types-FP/)
- [Recurrent Neural Networks lecture by Yoshua Bengio](http://videolectures.net/deeplearning2016_bengio_neural_networks/)
- [Practical Advice for Building Deep Neural Networks](https://pcc.cs.byu.edu/2017/10/02/practical-advice-for-building-deep-neural-networks/)