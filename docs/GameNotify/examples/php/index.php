<?php

function rsa_verify($data, $sign, $publicKey)
{
    $public_key = openssl_get_publickey($publicKey);
    if (empty($public_key)) {
        return false;
    }
    $sign = base64_decode($sign);
    $ok = openssl_verify($data, $sign, $public_key, OPENSSL_ALGO_SHA256);
    openssl_free_key($public_key);
    return $ok;
}


$data = "{\"Amount\":0.99,\"ExtraData\":\"{\\\"zoneId\\\":1000,\\\"gameGoodId\\\":\\\"4001\\\",\\\"money\\\":9900,\\\"roleName\\\":\\\"\\\",\\\"notifyUri\\\":\\\"http://localhost:8083/lt_charge\\\",\\\"roleId\\\":9192,\\\"extInfo\\\":2,\\\"goodGearId\\\":1,\\\"sdkProductId\\\":\\\"com.yostaren.revivedwitch.diamonds6\\\",\\\"productName\\\":\\\"6钻石\\\",\\\"orderId\\\":\\\"1229260561000\\\",\\\"ratio\\\":100}\",\"OrderID\":\"140088917161212164754\",\"ProductID\":\"com.yostaren.revivedwitch.diamonds6\",\"Type\":\"delivery\",\"UID\":\"1376172378933899192204\"}";
$sign = "QPsWdAq0Ywzh4CfvdUoCkZKsUCYfLIKuEhWZvymPm+ebU3QokcxWWN9JVzBuO92pob5qdqKGSkbvustnrNx5h39BNvXRaHRD5CuMlXNKG42vuzp+Dj7rVrIzhQPw8u8r4wvF1kRZ6FGbOWrqz9SsObvjQPKBCmtl7wNsuREEUfE=";
$publicKey = "-----BEGIN PUBLIC KEY-----公钥内容xxx-----END PUBLIC KEY-----\n";

$ok = rsa_verify($data, $sign, $publicKey);
echo $ok ? "Success" : "Fail";