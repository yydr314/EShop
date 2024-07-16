$(function () {
    loginApp.init();
})
var loginApp = {
    init: function () {
        this.getCaptcha()
        this.captchaImgChange()
    },
    getCaptcha: function () {
        //  使用隨機數避免瀏覽器緩存
        $.get("/admin/captcha?t=" + Math.random(), (res) => {
            console.log(res)
            $("#captchaId").val(res.captchaId)
            $("#captchaImg").attr("src", res.captchaImage)
        })
    },
    captchaImgChange: function () {
        var that = this;
        $("#captchaImg").click(() => {
            that.getCaptcha()
        })
    }
}