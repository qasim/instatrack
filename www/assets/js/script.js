window.minTagID = ''
var colorThief = new ColorThief()

$(document).ready(function () {
  var tag = getTag()
  if (tag != '') {
    var grid = $('.grid')
    for (i = 0; i < 20; i++) {
      grid.append('<div class="item index_' + i + '" id="index_' + i + '"></div>')
    }

    $("#tag").val(tag)
    getMedia(tag)
    var timer = setInterval(function () { getMedia(tag) }, 1400)
  }
})

function getTag() {
  var url = window.location.href
  if(url.indexOf('=') < 0) return ''
  return window.location.href.split('=')[1]
}

function getMedia(tag) {
  $.getJSON("/media?tag=" + tag + '&min_tag_id=' + window.minTagID, function (data) {
    if (data.data.length > 0) {
      window.minTagID = data.pagination.min_tag_id
    }
    var data = data.data
    var indices = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19]
    indices = shuffle(indices)
    data.forEach(function(e, i) {
      var item = $('#index_' + indices[i])
      var time = new Date(parseInt(data[i].created_time) * 1000).toISOString()
      var link = data[i].link
      var caption = data[i].caption.text
      preload([data[i].images.standard_resolution.url], function(url, c) {
        var rgb = c[0] + ',' + c[1] + ',' + c[2]
        item.css('z-index', '0')
        var newItem = $('<div class="item ' + item.attr('id') + '" style="display: none; z-index: 50" id="' + item.attr('id') + '" onclick="window.open(\'' + link + '\', \'_blank\')"><div class="info" title="' + time + '">' + caption + '</div></div>')
        newItem.insertBefore(item)
        newItem.css({
          'background': 'url(' + url + ') no-repeat center center'
        })
        document.styleSheets[0].addRule('#' + newItem.attr('id') + ':after', 'background-color: rgb(' + rgb + ')');
        $(".info").timeago()
        newItem.fadeIn(647, function () {
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
    cb(url, colorThief.getColor(img))
  }
  img.src = url
}
