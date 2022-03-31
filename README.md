# 一个分布式爬虫框架

使用 Golang 编写， JsonRPC 分布式通信，Elasticsearch 存储和搜索，Docker 部署，前端网页展示和分析的一个分布式爬虫框架。

## 构建

软件环境需求：

* Go >= 1.16
* Elasticsearch 6.x
* Redis >= 6.0

克隆代码到本地工作目录：

`$ git clone https://github.com/JohnGuoy/simple_httpd.git`

请在 config/config.ini 配置好 Elasticsearch 和 Redis 的 IP:Port。

## 运行

1 运行 persist/server/main.go 启动持久层的JsonRPC服务。

2 运行 worker/server/main.go 启动worker的JsonRPC服务。

3 运行 main.go 启动爬虫主程序爬取网络数据。

## 说明

当前程序默认爬取真爱网用户信息，用户可以开发定制自己想要爬取的网站，只需模仿 zhenai/parser/ 目录里的程序写网页数据解析器即可。

用户可以使用 Docker 来启动 Elasticsearch、Redis等组件，Redis用于URL去重。

当前程序前端展示页面暂缺，正在寻找好看的模板。当前程序有些代码文件尚未完成，正在编写中，估计还需一周时间才能全部完成。
