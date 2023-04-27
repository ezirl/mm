<html>
<head><title>{{.post.Title}}</title>
    <meta http-equiv="Refresh" content="0; URL={{.post.Link}}">
    <meta http-equiv="title" content="{{.post.Title}}">
    <meta http-equiv="description" content="{{.post.Description}}">

</head>
<body>
<div style="text-align: center; color: white;">
    <h1>Переадресация на страницу: <br><a href="{{.post.Link}}" style="color: #cc3333;">{{.post.Title}}</a></h1>
    {{.post.Link}}
</div>
<script type="text/javascript">
	(async (){
		await sleep(500);
	 window.location.href = "{{.post.Link}}"
	})
</script>
</body>
</html>
