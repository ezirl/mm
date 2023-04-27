function showTitle(title) {
  console.log(title);
}

function openPost(id) {
  $("#p" + id).toggle('normal');
  let arrow = $("#a" + id);
  if (arrow.html() === " ▼") {
    arrow.html(" ▲");
  } else {
    arrow.html(" ▼");
  }
}


/*
 * На сколько страниц автоматически прокручиваем до остановки
 */
var maxPages = 3;
/*
 * За сколько до конца страницы включаем прокрутку
 */
var scrollBufferRatio = 1 / 3;
/*
 * Где находится область для дополнения (содержимое)
 */
var elementContainerSelector = '#posts';
/*
 * Что делать если не найден новое содержимое
 */
let elementFindingErrorHandler = function() {
  throw "Could not find " + elementContainerSelector;
};
/*
 * Где находится выбор страниц для перехода
 */
let pagerSelector = '.pagination';
/*
 * Где найти ссылку на следующую страницу
 */

let nextPageSelector = '.pagination #next';

$elementContainer = $(elementContainerSelector);
$(pagerSelector).hide();

function load() {
  let $nextPage = $(nextPageSelector);
  // если следующей страницы нет...
  if ($nextPage.length === 0) {
    // то и загружать дальше нечего
    return true;
  }

  $.get($nextPage.attr('href'), function(data) {
    let $data = $(data);

    // найдем содержимое в следующей странице
    let $newContent = $data.find(elementContainerSelector);
    if ($newContent.length === 0) {
      // если элемент не найден, он может быть корневым
      $newContent = $data.filter(elementContainerSelector);
    }
    // если ничего не найдено, то это конкретно ошибка
    if ($newContent.length === 0) {
      elementFindingErrorHandler();
    }
    // содержимое из следующей страницы допишем
    // к содержимому страницы в окне
    $elementContainer.append($newContent.first().html());

    // найдем выбор страниц на следующей странице
    let $newPager = $data.find(pagerSelector);
    if ($newPager.length === 0) {
      // корневой элемент тоже поищем
      $newPager = $data.filter(pagerSelector);
    }
    // заменим выбор страниц новым
    let newPager = $newPager.first().html();
    // и скроем его пока не будет нужен
    $(pagerSelector).html(newPager).hide();

  })
}

$("ul.submenu").hover(function() {
  $(".link-lang").css("background-color", "transparent");
}, function() {
  $(".link-lang").css("background-color", "#cc3333");
});
$(".link-lang").hover(function() {
  $(this).css("background-color", "transparent");
}, function() {
  $(this).css("background-color", "#cc3333");
});
$(".weatherInformer21").css("width", "100%");

setTimeout(function() {
  $("#posts").prepend(
    "<div class='item' style='text-align: center; font-family: Times New Roman,sans-serif;'>" +
    "<a href='' style='display: block; font-size: 12px; padding: 3px;'>Читать еще...</a></div>"
  )
}, 1000 * 60); // время в мс

$(function() {
  $(window).scroll(function() {
    if ($(this).scrollTop() > 300) {
      $('#toTop').fadeIn();
    } else {
      $('#toTop').fadeOut();
    }
  });
  $('#toTop').click(function() {
    $('#toTop > span').empty();
    $('body,html').animate({
      scrollTop: 0
    }, 1500);

  });
});

$(".sidebar-item__head").click(function() {
  $(this).next('.sidebar-item__section').toggle('normal');
  if (($(this).find('#drop').html() === " ▼")) {
    $(this).find('#drop').html(" ▲");
  } else {
    $(this).find('#drop').html(" ▼");
  }
});
