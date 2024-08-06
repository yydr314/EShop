$(function () {
	baseApp.init();
})

var baseApp = {
	init: function () {
		this.initAside()
		this.confirmDelete()
		this.resizeIframe()
		this.changeStatus()
		this.changeNum()
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
	},
	changeNum: function () {
		$(".chSpanNum").click(function () {
			//	獲取element的數值
			var dataId = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var number = $(this).html().trim()

			var spanEl = $(this)

			//	建立一個input框，並且放入element中
			var input = $("<input style='width: 60px' value='' />")
			$(this).html(input)

			//	點擊後input框獲得焦點，並且將數值放入輸入框中
			$(this).trigger("focus").val(number)

			//	當元素已經有input時，再次點擊會再次觸發函式，這裡避免這個問題
			$(input).click(function (e) {
				e.stopPropagation()
			})

			//	滑鼠點擊input框外時給span賦值且發送ajax請求給後端
			$(input).blur(function () {
				var inputNum = $(this).val();
				spanEl.html(inputNum)

				$.get("/admin/changeNum", { id: dataId, table: table, field: field, num: inputNum },function(response){
					console.log(response)
				})
			})
		})
	}
}