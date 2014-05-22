var canvas = document.getElementById("game-board");
var canvasCtx = canvas.getContext("2d");
var boardTiles = 8;
var boardSpacing = 512 / boardTiles;

// --------------------------------------------------------------------
// Drawing code
// --------------------------------------------------------------------

for(var i = 0; i <= boardTiles; i++) {
	canvasCtx.moveTo(0, boardSpacing * i);
	canvasCtx.lineTo(512, boardSpacing * i);
	canvasCtx.moveTo(boardSpacing * i, 0);
	canvasCtx.lineTo(boardSpacing * i, 512);
}

canvasCtx.stroke();

function drawStone(player, x, y) {
	var centerX = x * boardSpacing + boardSpacing / 2;
	var centerY = y * boardSpacing + boardSpacing / 2;

	switch(player) {
		case '0': canvasCtx.fillStyle = '#FF0000'; break;
		case '1': canvasCtx.fillStyle = '#0000FF'; break;
	}

	canvasCtx.beginPath();
	canvasCtx.arc(centerX, centerY, boardSpacing * 0.4, 0, 2 * Math.PI, false);
	canvasCtx.fill();
	canvasCtx.stroke();
}

function setStatus(status) {
	$('#help').hide(400);
	$("#status").fadeOut(400, function() {
		$("#status").html(status).fadeIn();
	});
};

// --------------------------------------------------------------------
// UI code
// --------------------------------------------------------------------

canvas.addEventListener('click', function(ev) {
	var x = Math.floor((ev.pageX - canvas.offsetLeft) / boardSpacing);
	var y = Math.floor((ev.pageY - canvas.offsetTop) / boardSpacing);
	placeStone(x, y);
}, false);

// --------------------------------------------------------------------
// Netcode
// --------------------------------------------------------------------

var playerId;

function placeStone(x, y) {
	var url = '/place?g=' + gameKey + "&p=" + playerToken + "&x=" + x + "&y=" + y;
	var xhr = new XMLHttpRequest();

	xhr.open('POST', url, true);
	xhr.send();
}

var chan = new goog.appengine.Channel(playerToken);
var sock = chan.open();

sock.onclose = function() {
	setStatus("Session terminated");
};

sock.onerror = function(error) {
	setStatus("Error: " + JSON.stringify(error));
};

sock.onmessage = function(data) {
	split = data.data.split(":");

	switch(split[0]) {
		case "partner_dc":
			setStatus("Partner DC");
			break;
		case "begin":
			setStatus("Your turn");
			playerId = "0";
			break;
		case "wait":
			setStatus("Opponent's turn");
			playerId = "1";
			break;
		case "move":
			drawStone(split[1], split[2], split[3]);

			if(split[1] == playerId) {				
				setStatus("Opponent's turn");
			} else {
				setStatus("Your turn");
			}

			break;
		default:
			alert(JSON.stringify(data));
			break;
	}
};
