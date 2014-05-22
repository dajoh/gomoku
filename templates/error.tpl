<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<meta name="description" content="Gomoku" />
		<meta name="author" content="Daniel Johannesson Lindberg, Leon Landvall" />
		<title>Gomoku! - Error</title>
		<link href="/css/bootstrap.min.css" rel="stylesheet" />
		<link href="/css/bootstrap-theme.min.css" rel="stylesheet" />
		<!--[if lt IE 9]>
			<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
			<script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
		<![endif]-->
	</head>
	<body>
		<div class="container">
			<div class="page-header">
				<h1>Gomoku!</h1>
			</div>
			<div class="alert alert-danger">
				<h4>Error!</h4>
				{{.}}
			</div>
			<hr />
			<span>No rights reserved.</span>
		</div>
		<script src="/js/jquery.min.js"></script>
		<script src="/js/bootstrap.min.js"></script>
	</body>
</html>
