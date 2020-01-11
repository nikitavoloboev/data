# Neural networks

## Notes

- Neural Networks are great identifying patterns in data. As a classic example, if you wanted to predict housing prices, you could build a data set that maps features about houses (square feet, location, proximity to Caltrain, etc) onto their actual price, and then train a network to recognize the complex relationship between features and pricing. Training happens by feeding the network features, letting it make a guess about the price, and then correcting the guess (backpropagation).
  - Convolutional Neural Networks work similarly, but with images. Instead of giving a CNN discrete features, you'll usually just use the pixels of the image itself. Through a series of layers, the CNN is able to build features itself (traditionally things like edges, corners) and learn patterns in image data. For example, a CNN might be trained on a dataset that maps images onto labels, and learn how to label new images on its own.

## Links

- [But what is a Neural Network? | Deep learning, chapter 1 (2017)](https://www.youtube.com/watch?v=aircAruvnKk)
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
- [ONNX](https://github.com/onnx/onnx) - Open Neural Network Exchange. ([Scoring ONNX ML Models with Scala](https://www.youtube.com/watch?v=HyYpMJNVoVk))
- [DyNet](https://github.com/clab/dynet) - Dynamic Neural Network Toolkit.
- [gonn](https://github.com/sausheong/gonn) - Building a simple neural network in Go.
- [Neural Ordinary Differential Equations (2018)](https://arxiv.org/abs/1806.07366) - [Video explanation](https://www.youtube.com/watch?v=AD3K8j12EIE) | [Notes](https://github.com/llSourcell/Neural_Differential_Equations/blob/master/Neural_Ordinary_Differential_Equations.ipynb)
- [MindsDB](https://github.com/mindsdb/mindsdb) - Framework to streamline use of neural networks.
- [Neural Network framework in 25 LOC](https://gist.github.com/macournoyer/620a8ba4a2ecd6d6feaf)
- [Learning and Processing over Networks (2019)](https://github.com/rodrigo-pena/amld2019-graph-workshop) - Workshop presented by Michaël Defferrard and Rodrigo Pena at the Applied Machine Learning Days in January 2019.
- [The Next Generation of Neural Networks by Geoffrey Hinton (2007)](https://www.youtube.com/watch?v=AyzOUbkUf3M)
- [Who Invented Backpropagation? (2014)](http://people.idsia.ch/~juergen/who-invented-backpropagation.html)
- [Deep Learning in Neural Networks: An Overview (2015)](http://people.idsia.ch/~juergen/deep-learning-overview.html)
- [Neural Networks (E01: introduction) (2018)](https://www.youtube.com/watch?v=bVQUSndDllU) - This series is intended as a light introduction to neural networks, with a focus on the task of classifying handwritten digits.
- [Machine Learning for Beginners: An Introduction to Neural Networks (2019)](https://victorzhou.com/blog/intro-to-neural-networks/)
- [A Recipe for Training Neural Networks (2019)](https://karpathy.github.io/2019/04/25/recipe/)
- [Exploring Neural Networks with Activation Atlases (2019)](https://distill.pub/2019/activation-atlas/)
- [Curated list of neural architecture search and related resources](https://github.com/D-X-Y/Awesome-NAS#readme)
- [Weight Agnostic Neural Networks (2019)](https://weightagnostic.github.io/) ([HN](https://news.ycombinator.com/item?id=20160693))
- [Geoffrey Hinton explains the evolution of neural networks (2019)](https://www.wired.com/story/ai-pioneer-explains-evolution-neural-networks/)
- [Evolved Turing neural networks](http://compucology.net/evolved)
- [Intelligent Machinery paper by Alan Turing](https://weightagnostic.github.io/papers/turing1948.pdf)
- [SRU](https://github.com/taolei87/sru) - Training RNNs as Fast as CNNs.
- [ODIN](https://github.com/facebookresearch/odin) - Out-of-Distribution Detector for Neural Networks.
- [Ask HN: What Neural Networks/Deep Learning Books Should I Read? (2019)](https://news.ycombinator.com/item?id=20674745)
- [Python vs Rust for Neural Networks (2019)](https://ngoldbaum.github.io/posts/python-vs-rust-nn/) ([HN](https://news.ycombinator.com/item?id=20728288))
- [Exploring Weight Agnostic Neural Networks (2019)](https://ai.googleblog.com/2019/08/exploring-weight-agnostic-neural.html) ([HN](https://news.ycombinator.com/item?id=20817083))
- [Neural Networks, Types, and Functional Programming (2015)](https://colah.github.io/posts/2015-09-NN-Types-FP/)
- [Glow](https://github.com/pytorch/glow) - Compiler for Neural Network hardware accelerators.
- [Go Neural Net Sandbox](https://github.com/lightvector/GoNN) - Sandbox for personal experimentation in Go neural net training and Go AI.
- [layer](https://github.com/cloudkj/layer) - Neural network inference the Unix way.
- [XNNPACK](https://github.com/google/XNNPACK) - Highly optimized library of floating-point neural network inference operators for ARM, WebAssembly, and x86 (SSE2 level) platforms.
- [LSTM implementation explained (2015)](http://apaszke.github.io/lstm-explained.html)
- [The Neural Process Family](https://github.com/deepmind/neural-processes) - Contains notebook implementations of the following Neural Process variants: Conditional Neural Processes (CNPs), Neural Processes (NPs), Attentive Neural Processes (ANPs).
- [Notes on Neural Nets](https://wiki.kourouklides.com/wiki/Artificial_Neural_Network)
- [Graph Neural Tangent Kernel: Fusing Graph Neural Networks with Graph Kernels (2019)](https://github.com/KangchengHou/gntk)
- [RNNoise](https://github.com/xiph/rnnoise) - Recurrent neural network for audio noise reduction.
- [Hacking Neural Networks](https://github.com/Kayzaks/HackingNeuralNetworks) - Short introduction on methods that use neural networks in an offensive manner.
- [Distilling knowledge from Neural Networks to build smaller and faster models (2019)](https://blog.floydhub.com/knowledge-distillation/)
- [Neural Network Processing Neural Networks: An efficient way to learn higher order functions (2019)](https://arxiv.org/abs/1911.05640)
- [Building a neural net from scratch in Go (2017)](https://datadan.io/neural-net-with-go)
- [SparseConvNet](https://github.com/btgraham/SparseConvNet) - Spatially-sparse convolutional neural network.
- [Norse](https://github.com/electronicvisions/norse) - Library to do deep learning with spiking neural networks.
- [Tensor Programs I: Wide Feedforward or Recurrent Neural Networks of Any Architecture are Gaussian Processes (2019)](https://arxiv.org/abs/1910.12478)
- [Single Headed Attention RNN: Stop Thinking With Your Head (2019)](https://arxiv.org/abs/1911.11423) ([HN](https://news.ycombinator.com/item?id=21647804))
- [Visualizing the Loss Landscape of Neural Nets](https://github.com/tomgoldstein/loss-landscape)
- [primitiv](https://github.com/primitiv/primitiv) - Neural Network Toolkit.
- [On the Relationship between Self-Attention and Convolutional Layers (2019)](https://openreview.net/forum?id=HJlnC1rKPB) ([Code](https://github.com/epfml/attention-cnn))
