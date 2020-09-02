你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# ENV to ECS Tool

This simple go script will ingest an environment definitions file and convert it into a JSON blob that ECS can consume in a task definition.

## Basic Usage

```bash
./env-to-ecs -i <INFILE_TO_PARSE> -o <OUTFILE_TO_WRITE_TO> 
```

Where an input looks like:

```dotenv
A=B
```

An the output will look like:

```json
[{"name":"A","value":"B"}]
```

### Features

* -i | --infile: pass an infile in dotenv format that will be parsed
    * NOTE: the tool will raise an error and exit if you do not set this arg
* -o | --outfile: pass an outfile in JSON format that will be written to
    * NOTE: If you do not set this arg, the output will be written to stdout.
    * NOTE: if you pass an outfile that does not exist, the outfile will be created for you
* -v | --variable: pass in extra key=value pairs one at a time.
    * NOTE: Must be in dotenv format ie. A=B
    * NOTE: The arg can be passed multiple times
* Supports comments in the infile ie. `# this is a comment`. These will be parsed out.
    
    
#### Why does this tool exist?

We wanted a way to configure our environments in one place, and know that those configuration would be propagated elsewhere automagically.
For us, an environment definitions file is that single source of truth, and all other configuration files should be dependant on that.
This will hopefully lead to less configuration drift and questions down the road.

#### Developing

Simplest way to get started is to clone the repo and run the `./ci/scripts/ensure_deps.sh`. We use glide for dependencies, so make sure you have that installed. Once you have that running the scripts in that directory should help. They're hopefully named in an intuitive way.
