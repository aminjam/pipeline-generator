# Concourse Pipeline Generator
I wanted to be able to define a go template, so that we could define a `.Range` for our repeatable resrouces and jobs. If you are looking for a more robust solution without regex replacement [checkout goflat](https://github.com/aminjam/goflat).

# How to use
```
git clone git@github.com:aminjam/pipeline-generator.git
# assuming you have yaml.v2 or go get it
go get gopkg.in/yaml.v2
./generator.go > /tmp/final-pipeline.yml
```

# Modify to your needs
* Change `data` and `Structure` in `generator.go` depending on your needs
* Change `./fixtures/pipeline.yml` to use your own pipeline template

# Hopes and Wishes
* Add this functionality to the core [fly](https://github.com/concourse/fly/blob/master/template/variables.go) cli
