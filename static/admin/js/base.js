$(function () {
	baseApp.init();
})

var baseApp = {
	init: function () {
		this.initAside()
		this.confirmDelete()
		this.resizeIframe()
		this.changeStatus()
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
	},
	changeStatus: function () {
		$(".chStatus").click(function () {
			var dataId = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var el = $(this)
			$.get("/admin/changeStatus", { id: dataId, table: table, field: field }, function (response) {
				console.log(response)
				if (response.success) {
					if (el.attr("src").indexOf("yes") != -1) {
						el.attr("src", "/static/admin/images/no.gif")
					} else {
						el.attr("src", "/static/admin/images/yes.gif")
					}
				}
			})
		})
	}
}