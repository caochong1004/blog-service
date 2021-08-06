package errcode

var (
	ErrorGetTagFail = NewError(20010000, "获取单个标签失败")
	ErrorGetTagListFail = NewError(20010001, "获取列表失败")
	ErrorCreateTagFail = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail = NewError(20010004, "删除标签失败")
	ErrorCountTagFail = NewError(20010005, "统计标签失败")
	ErrorGetArticleFail = NewError(20010006, "获取单个文章失败")
	ErrorGetArticleListFail = NewError(20010007, "获取文章列表失败")
	ErrorCreateArticleFail = NewError(20010008, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20010009, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20010010 ,"删除文章失败")
	ErrorUploadFileFail = NewError(20030001, "上传文件错误")

	)
