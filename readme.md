## 如何使用
先从 releases 中下载适用于自己系统的可执行文件

直接运行将读取同目录下“config.yml”作为配置文件，参考 [config.yml](https://github.com/realZnS/jxnu-srun-go/blob/master/config.yml)
```
./srun
```
当然，也可以指定配置文件：
```
./srun -c /path/to/example.yml
```
在命令中加入 `-v` 将展示更多信息用于 debug。
## 自编译
```
git clone https://github.com/realzns/jxnu-srun-go.git && cd jxnu-srun-go
chmod +x build.sh
./build.sh
```
编译的可执行文件将存放在 ./out 目录下
## 免责声明
本项目仅供个人学习使用，请勿用于违反校规之处。
## 致谢
本项目是 [jxnu_srun](https://github.com/realzns/jxnu_srun) 的 Golang 实现，感谢 [@huxiaofan](https://github.com/huxiaofan1223) 的原程序和博客分析。
