$(function () {
	baseApp.init();
})

var baseApp = {
	init: function () {
		this.initAside()
	},
	initAside: function () {
		$('.aside h4').click(function () {
			$(this).siblings('ul').slideToggle();
		})
	},
	resizeIframe: function () {
		$("#rightMain").height($(window).height() - 80)
	},
	confirmDelete: function () {
		$(".delete").click(function () {
			var flag = confirm("您確定要刪除嗎？")
			return flag
		})
	}
}