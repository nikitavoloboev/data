# Python
Currently using more and more of this language.

```python{5-8}
import torch

def encode(points, padded_length=100):
    input_tensor = torch.zeros([2, padded_length])
    for i in range(min(padded_length, len(points))):
        input_tensor[0][i] = points[i][0] * 1.0 # cast to float
        input_tensor[1][i] = points[i][1] * 1.0
        continue
    return input_tensor
```