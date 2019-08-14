[![Build Status](https://travis-ci.org/rc452860/vnet.svg?branch=master)](https://travis-ci.org/rc452860/vnet)

# 前端
[https://github.com/jarvanh/SSRPanel](https://github.com/jarvanh/SSRPanel)

## 运行
```
wget https://github.com/jarvanh/vnet/releases/download/0.0.7/vnet_linux_arm64 -O vnet && chmod +x vnet && ./vnet

配置好数据库后按ctrl + c退出使用nohup启动
nohup ./vnet>vnet.log 2>&1 &
```
重新启动
```
kill -9 $(ps aux | grep '[v]net' | awk '{print $2}') && nohup ./vnet>vnet.log 2>&1 &
```

## 支持加密方式
```
aes-256-cfb
bf-cfb
chacha20
chacha20-ietf
aes-128-cfb
aes-192-cfb
aes-128-ctr
aes-192-ctr
aes-256-ctr
cast5-cfb
des-cfb
rc4-md5
salsa20
aes-256-gcm
aes-192-gcm
aes-128-gcm
chacha20-ietf-poly1305
```
## 注意事项
config.json配置文件中的所有时间单位都为毫秒

升级后需要删除原有config.json重新生成