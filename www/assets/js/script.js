window.minTagID = ''

$(document).ready(function () {
  var tag = getTag()
  if (tag != '') {
    var grid = $('.grid')
    for (i = 0; i < 20; i++) {
      grid.append('\
<div class="item" id="index_' + i + '"> \
  <div class="info"> \
    <div class="photo"></div> \
    <div class="text"></div> \
  </div> \
</div>')
    }

    $("#tag").val(tag)
    getMedia(tag)
    // var timer = setInterval(function () { getMedia(tag) }, 2000)
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
    console.log(data)
    for(i = 0; i < data.length; i++) {
      var item = $('#index_' + i)
      item.css({
        'background-image': 'url(' + data[i].images.standard_resolution.url + ')'
      })
      item.find('.photo').css({
        'background-image': 'url(' + data[i].user.profile_picture + ')'
      })
      item.find('.text').html(data[i].caption.text)
    }
  })
}
