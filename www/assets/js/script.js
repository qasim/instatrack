$(document).ready(function () {
  window.minTagID = ''
  window.grid = $('.grid')
  window.indices = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19]
  window.direction = ['left', 'right', 'up', 'down']
  window.inactive = null
  var tag = getTag()
  if (tag != '') {
    for (i = 0; i < 20; i++) grid.append('<div class="item index_' + i + '" id="index_' + i + '"></div>')
    $('#tag').val(tag)
    getMedia(tag)
    var timer = setInterval(function () { getMedia(tag) }, 1400)
  }

  $(document).mousemove(function() {
      clearTimeout(window.inactive)
      $('.track-box').show()
      window.inactive = setTimeout(function () {
          $('.track-box').hide()
      }, 10000)
  }).mouseleave(function() {
      clearTimeout(window.inactive)
      $('.track-box').hide()
  })
})

function getTag() {
  var url = window.location.href
  if(url.indexOf('=') < 0) return ''
  return window.location.href.split('=')[1]
}

function getMedia(tag) {
  $.getJSON('/media?tag=' + tag + '&min_tag_id=' + window.minTagID, function (data) {
    if (data.data.length > 0) {
      window.minTagID = data.pagination.min_tag_id
    }
    var data = data.data
    window.indices = shuffle(window.indices)
    data.forEach(function(e, i) {
      var item = $('#index_' + window.indices[i])
      var time = new Date(parseInt(data[i].created_time) * 1000).toISOString()
      var link = data[i].link
      var caption = data[i].caption.text
      preload([data[i].images.standard_resolution.url], function(url, c) {
        var rgb = c[0] + ',' + c[1] + ',' + c[2]
        item.css('z-index', '0')
        window.grid.append('<div class="item ' + item.attr('id') + ' ' + data[i].id + '" style="display: none; z-index: 50; background: url(' + url + ') no-repeat center center" id="' + item.attr('id') + '" onclick="window.open(\'' + link + '\', \'_blank\')"><div class="color" style="background-color: rgb(' + rgb + ')"></div></div>')
        var newItem = $('.' + data[i].id)
        var randomDirection = window.direction[Math.floor(Math.random()*4)]
        newItem.show('slide', {
          direction: randomDirection,
          easing: 'easeOutExpo'
        }, 1200, function() {
          item.remove()
        })
      })
    })
  })
}

function shuffle(array) {
  var currentIndex = array.length, temporaryValue, randomIndex
  while (0 !== currentIndex) {
    randomIndex = Math.floor(Math.random() * currentIndex)
    currentIndex -= 1
    temporaryValue = array[currentIndex]
    array[currentIndex] = array[randomIndex]
    array[randomIndex] = temporaryValue
  }
  return array
}

function preload(url, cb) {
  var img = new Image()
  img.crossOrigin = 'Anonymous'
  img.onload = function() {
    var vibrant = new Vibrant(img)
    var swatches = vibrant.swatches()
    var color = swatches['LightMuted']
    if (!color) {
      color = swatches['Muted']
    }
    if (!color) {
      color = swatches['DarkMuted']
    }
    cb(url, color.getRgb())
  }
  img.src = url
}
