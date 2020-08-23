---
title: "关于对扫描器的一些想法及实现"
date: 2020-08-23T15:55:19+08:00
draft: true
---

# [模板参考至Xray](https://docs.xray.cool/#/guide/poc)
# [github地址](https://github.com/lovetrap)

### POC检测模板解析
#### yaml -> sets -> Rules
```
{{ set golbal var}} ==> regexp replace ==> value 
==> {{.*?}} ==> value ==> map[value]

==> {{{r1}} * {{r2}}} ==> {.*} ==> 同上 ==> result * result
```

#### yaml -> expression

```
/* 关于解析cel伪代码，这是我的实现方法 */
{
response.body ==> []byte
response.headers ==> map
response.content_type ==> []string
} ==> set

{
response.body.bcontains(b"??")
==> bcontains(b"??", response.body) ==> bool
}

{
response.headers ==> Type ==> decls.Map
response.headers["??"] == "??" ==> bool

response.headers["??"].contains("??")
==> contains("??", response.headers) ==> bool
}

{
"??".bmatches(response.body)
==> bmatches(response.body, "??") ==> bool
}
```

### 对于模板解析我增加了一些自定义检测模板
#### 例如对目录进行扫描，探测是否有信息泄露或者其他
```
[module demo]
name: dirscan-yaml-backstage
wordlist: .\wordlist\test.txt
search: (?P<username>(账[户号]|管理员账[户号]|username|password|密码))

```
#### 下面是解析后的伪代码
```
if regexp(search, html) != nil ==> 输出信息

```
### 下面是我对上面实现的demo
 ![](/img/001.png)
