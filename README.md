

#  Gee

golang web 框架

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
    
<h3 align="center">Gee</h3>
  <p align="center">
    "Gee" web 框架
    <br />
    <a href="https://github.com/learnselfs/gee"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com/learnselfs/gee">查看Demo</a>
    ·
    <a href="https://github.com/learnselfs/gee/issues">报告Bug</a>
    ·
    <a href="https://github.com/learnselfs/gee/issues">提出新特性</a>
  </p>

</p>


本篇README.md面向开发者

## 目录

- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
    - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
    - [版本更新](#版本更新)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南


###### 开发前的配置要求

1. go mode init project

###### **安装步骤**

1. go get https://github.com/learnselfs/gee

### 文件目录说明
eg:
```
├─config        # 配置项
├─core          # 核心功能
├─doc           # 文档
├─middleware    # 中间件文件
├─test          # 测试和示例
└─utils         # 工具函数
```

### 开发的架构
请阅读[net/http](https://github.com/golang/go/tree/master/src/net/http) 查阅为该项目的架构。

### 部署
暂无

### 使用到的框架


### 贡献者
请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目
贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制
该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

#### 版本更新
- v1.0
  - Complete basic utils log settings
  - Complete basic http Handler interface
- v2.0
  - Complete basic request response context management
- v3.0
  - Complete basic router and handler binding mapping
  - Optimize router trie configuration
- v4.0
  - complete basic route group
- v5.0
  - Complete basic middleware
- v6.0
  - Complete basic html template

### 版权说明
该项目签署了 apache 授权许可，详情请参阅 [LICENSE.txt](https://github.com//learnselfs/gee/blob/main/LICENSE)

### 鸣谢
- [Gin](https://github.com/gin-gonic/gin)
- [geektutu](https://github.com/geektutu/7days-golang)

<!-- links -->
[your-project-path]: learnselfs/gee
[contributors-shield]: https://img.shields.io/github/contributors/learnselfs/gee?style=flat-square
[contributors-url]: https://github.com/learnselfs/gee/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/learnselfs/gee.svg?style=flat-square
[forks-url]: https://github.com/learnselfs/gee/network/members
[stars-shield]: https://img.shields.io/github/stars/learnselfs/gee.svg?style=flat-square
[stars-url]: https://github.com/learnselfs/gee/stargazers
[issues-shield]: https://img.shields.io/github/issues/learnselfs/gee.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/learnselfs/gee.svg
[license-shield]: https://img.shields.io/github/license/learnselfs/gee.svg?style=flat-square
[license-url]: https://github.com/learnselfs/gee/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian