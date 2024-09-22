```
source /opt/homebrew/Caskroom/miniconda/base/etc/fish/conf.d/conda.fish
conda activate mlx_code
```

- install latest python version inside `conda env`
```
conda install python=3.10
```

```
# create env
conda create -n python python=3.12
# activate
conda activate python
# install packages
conda install pymc


# can also

# change python inside running conda env
conda install python=3.12
```