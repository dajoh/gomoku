<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<meta name="description" content="Gomoku" />
		<meta name="author" content="Daniel Johannesson Lindberg, Leon Landvall" />
		<title>Gomoku! - In Game</title>
		<link href="/css/bootstrap.min.css" rel="stylesheet" />
		<link href="/css/bootstrap-theme.min.css" rel="stylesheet" />
		<!--[if lt IE 9]>
			<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
			<script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
		<![endif]-->
		<script>
			gameKey = window.location.search.split('=')[1];
			playerToken = "{{.}}";
		</script>
	</head>
	<body>
		<div class="container">
			<div class="page-header">
				<h1>Gomoku! <small id="status">Waiting for second player</small></h1>
			</div>
			<div id="help" class="alert alert-info">
				Link this page to a friend to play!
			</div>
			<div id="winner" class="alert alert-success" style="display: none">
				You won!
			</div>
			<div id="loser" class="alert alert-danger" style="display: none">
				You lost!
			</div>
			<p>
				The game of Gomoku is simple, just place 5 stones in a row to win.
			</p>
			<canvas id="game-board" width="512" height="512">
			</canvas>
			<hr />
			<span>No rights reserved.</span>
		</div>
		<script src="/_ah/channel/jsapi"></script>
		<script src="/js/jquery.min.js"></script>
		<script src="/js/bootstrap.min.js"></script>
		<script src="/js/gomoku.js"></script>
	</body>
</html>
