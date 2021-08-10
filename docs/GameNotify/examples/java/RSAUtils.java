package com.yostar.rsatest;

import android.util.Base64;

import java.security.KeyFactory;
import java.security.PublicKey;
import java.security.Signature;
import java.security.spec.X509EncodedKeySpec;

public class RSAUtils {
    public static void main() {
        String content = "{\"Amount\":0.99,\"ExtraData\":\"{\\\"OrderNo:\\\":\\\"TPS-af88932e-f994-41f4-8a93-460742983c3f\\\"}\",\"OrderID\":\"2743461520780335155717\",\"ProductID\":\"com.yostar.arknights.originiteprime1\",\"Type\":\"delivery\",\"UID\":\"7265561517399945859280\"}";
        String sign = "NRI17G+U0k+FRHzBxDsprEX01UhR6GeWlyT9F58heMe8K1kR5AjKd2oxOjJxlzqNm27jfkAAEZno+3p3n0bbsH7/A7F1mb3fzTBu/ZSNEErmmq1rJZ77KIGPvnyZpTKh9DgGtMAV4YIb1gc2fdfBlfJTZKpA6OmesLNo/PIFH14=";
        RSAUtils.checkSign(content, sign);
    }
    /**
     * RSA验签名检查 SHA256 加密 ，数据都需要经过base64解密
     * @param content 待签名数据,即服务器的返回数据
     * @param sign 签名值，即服务器签名处理后的sign值
     * @return 布尔值
     */
    public static boolean checkSign(String content, String sign){
        String publicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZKvOO0FTLLfNEJoCCHGxJr0Bx\n" +
                "你的公钥xxx" +
                "rI5VLEIvUYNmS0rp5wIDAQAB";
        try {
            KeyFactory keyFactory = KeyFactory.getInstance("RSA");
            X509EncodedKeySpec keySpec = new X509EncodedKeySpec(Base64.decode(publicKey.getBytes(), Base64.DEFAULT));
            PublicKey publicK = keyFactory.generatePublic(keySpec);

            Signature signature = Signature.getInstance("SHA256WithRSA");
            signature.initVerify(publicK);
            signature.update(content.getBytes());
            return signature.verify(Base64.decode(sign.getBytes(), Base64.DEFAULT));
        } catch (Exception e) {
            e.printStackTrace();
            return false;
        }
    }

}
