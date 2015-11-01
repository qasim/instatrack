if(window.location.href.match("tag=")) {
  document.getElementById('home').classList.add('hidden')
  document.getElementById('grid').classList.remove('hidden')

  console.log('Making socket connection...')
  var socket = new WebSocket("ws://localhost:3000/socket")

  socket.onopen = function() {
    console.log('Opened')
    socket.send("Connection init")
  }

  socket.onmessage = function(e) {
    console.log("Received: " + e.data)
  }
}
