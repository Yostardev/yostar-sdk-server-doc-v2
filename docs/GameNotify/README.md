[TOC]

# 游戏服接入 - 说明文档

# 1. 支付/退款通知

## 1.1 通知规则

**规则**

- 支付完成后，SDK会把相关支付结果和订单信息发送给游戏服，游戏服需要接收处理该消息，并返回应答。

- 如果SDK收到游戏服的应答不符合规范或超时(10s)，SDK认为通知失败。

- SDK会通过一定的策略定期重新发起通知，尽可能提高通知的成功率，但SDK不保证通知最终能成功。

- 通知频率为`0s x3` / `15s x2` / `30s` / `3m` / `10m` / `20m` /` 30m x3` / `60m` / `3h x3` / `6h x2` - 总计 `24h4m`

**提醒**

- 同样的通知可能会多次发送给游戏服。游戏服必须能够正确处理重复的通知。 当游戏服收到通知进行处理时，先检查对应业务数据的状态，并判断该通知是否已经处理。如果未处理，则再进行处理；如果已处理，则直接返回结果成功。
- 在对业务数据进行状态检查和处理之前，要采用数据锁进行并发控制，以避免函数重入造成的数据混乱。

- 游戏服对于支付成功结果通知的内容一定要做签名验证，并校验通知的信息是否与游戏服侧的信息一致，防止数据泄漏导致出现“假通知”，造成资金损失

## 1.2 接口链接

该链接是通过[SDK后台管理] 服务端配置的`Pay.Tags`设置，如果链接无法访问，游戏服将无法接收到SDK通知。

通知url必须为直接可访问的url，不能携带参数。示例： https://test-sdk.com/pay/notify

## 1.3 通知报文

通知是以`POST`方法访问，通知的数据以`application/json`格式通过请求主体（BODY）传输。

出于安全考虑，SDK会对数据进行RSA签名。游戏服需要先对通知数据进行RSA验签，进而判断通知的有效性。

## 1.4 通知示例

```json
{
    "Data": "{\"Amount\":0.99,\"ExtraData\":\"{\\\"OrderNo:\\\":\\\"123456789\\\"}\",\"OrderID\":\"5ff8282bc5306f9146884389\",\"ProductID\":\"112334\",\"Type\":\"delivery\",\"UID\":\"5fec46083d81a400012b38b7\"}",
    "Sign": "gBoVBu9C6CeU/hsEA7k3/CqSxb8JEGeiE8zorgg3nfWoC2/0lXvBGLPJuRU4YNZSVaQS38wKEQfsXJEPOXAp68AnLmqjvPPsvg/fWiN3phcX/ac9KtX/VaIKq+zMRqFUw2mMFNnbI4Y5V5RuIUY9jefd/hCpOFOE2cKeXzaDltg="
}
```

## 1.5 通知参数

| 变量名 |  类型  | 描述         |
| ------ | :----: | ------------ |
| Data   | string | 签名原始数据 |
| Sign   | string | RSA签名结果  |

**Data字段解析成json后参数含义**

| 变量名    | 类型    | 描述                                                         |
| :-------- | :------ | ------------------------------------------------------------ |
| Type      | string  | 通知类型<br/>`delivery`为下单付款成功通知<br/>`refund`为用户退款成功通知 |
| Amount    | float64 | 用户付款金额                                                 |
| ExtraData | string  | 游戏服发起订单时,传入的数据                                  |
| ProductID | string  | 游戏服商品ID                                                 |
| OrderID   | string  | SDK订单号                                                    |
| UID       | string  | SDK用户ID                                                    |

## 1.6 验签步骤

### 1.6.1 获取RSA公钥

> RSA public key 由 SDK服务器 提供

### 1.6.2 调用示例

- [Golang](examples/golang/main.go)
- [Java](examples/java/RSAUtils.java)
- [PHP](examples/php/index.php)

## 1.7 通知应答

游戏服后台在正确处理回调之后，需要在10s内返回`200`或者`204`的HTTP状态码。

其他的状态码，SDK支付均认为通知失败，并按照前述的策略定期发起通知。

注意，当游戏服后台应答失败时，SDK支付将记录下应答的报文，建议游戏服按照以下格式返回。

```json
{
    "Code": "ERROR_NAME",
    "Msg": "ERROR_DESCRIPTION"
}
```

# 2. 验证用户信息接口

## 2.1 接口说明

> 此接口用于CP验证用户信息是否存在

## 2.2 接口地址

- 测试环境`https://test-sdk-api.yostarplat.com/game/check-user`
- 预发布日服`https://staging-jp-sdk-api.yostarplat.com/game/check-user`
- 预发布美服`https://staging-en-sdk-api.yostarplat.com/game/check-user`
- 正式环境日服`https://jp-sdk-api.yostarplat.com/game/check-user`
- 正式环境美服`https://en-sdk-api.yostarplat.com/game/check-user`

## 2.3 接口参数

> Content-Type为application/json

| 参数  | 示例                                     | 说明          |
| ----- | ---------------------------------------- | ------------- |
| PID   | KR-BiLanHangXian                         | 项目ID/游戏ID |
| UID   | 63052536224721998718                     | 用户编号      |
| Token | 08cf45903c141761854a76afeeed92693626b2e3 | 登录令牌      |

## 2.4 请求示例

```json
{
    "PID": "KR-BiLanHangXian",
    "UID": "39710899001880576",
    "Token": "08cf45903c141761854a76afeeed92693626b2e3"
}
```

## 2.5 响应说明

- 成功

> HTTP Code 为 200时,表示用户存在令牌匹配正确,响应报文如下

```json
{
    "Code": 200,
    "Data": {},
    "Msg": "OK"
}
```

- 失败

> HTTP Code 为404时,表示用户不存在或令牌不匹配,响应报文如下

```json
{
    "Code": 404,
    "Data": {},
    "Msg": "User does not exist"
}
```



