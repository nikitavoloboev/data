# Neural networks

[This](https://www.youtube.com/watch?v=aircAruvnKk) is a great introduction.

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
- [Differentiable Architecture Search](https://github.com/quark0/darts) - Code for [DARTS: Differentiable Architecture Search](https://arxiv.org/abs/1806.09055) paper.
- [TensorSpace.js](https://github.com/tensorspace-team/tensorspace) - Neural network 3D visualization framework, build interactive and intuitive model in browsers, support pre-trained deep learning models from TensorFlow, Keras, TensorFlow.js
- [UIS-RNN](https://github.com/google/uis-rnn) - Library for the Unbounded Interleaved-State Recurrent Neural Network (UIS-RNN) algorithm, corresponding to the paper Fully Supervised Speaker Diarization.
- [ONNX](https://github.com/onnx/onnx) - Open Neural Network Exchange.
- [DyNet](https://github.com/clab/dynet) - Dynamic Neural Network Toolkit.
- [gonn](https://github.com/sausheong/gonn) - Building a simple neural network in Go.
- [Neural Ordinary Differential Equations (2018)](https://arxiv.org/abs/1806.07366) - [Video explanation](https://www.youtube.com/watch?v=AD3K8j12EIE) | [Notes](https://github.com/llSourcell/Neural_Differential_Equations/blob/master/Neural_Ordinary_Differential_Equations.ipynb)
- [MindsDB](https://github.com/mindsdb/mindsdb) - Framework to streamline use of neural networks.
- [Neural Network framework in 25 LOC](https://gist.github.com/macournoyer/620a8ba4a2ecd6d6feaf)
- [Learning and Processing over Networks (2019)](https://github.com/rodrigo-pena/amld2019-graph-workshop) - Workshop presented by Michaël Defferrard and Rodrigo Pena at the Applied Machine Learning Days in January 2019.
- [The Next Generation of Neural Networks by Geoffrey Hinton (2007)](https://www.youtube.com/watch?v=AyzOUbkUf3M)
- [Who Invented Backpropagation? (2014)](http://people.idsia.ch/~juergen/who-invented-backpropagation.html)
- [Deep Learning in Neural Networks: An Overview (2015)](http://people.idsia.ch/~juergen/deep-learning-overview.html)
