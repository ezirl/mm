<div class="content-header blue">
    {{if .site}}
        Новости площадки - {{.site.Name}}:
    {{else}}
        Последние новости{{end}}</div>
<div id="posts">
    {{ range $item := .posts }}

        <div class="item" data-title="{{$item.Title}}">
            <span class="date"> {{ dateformat $item.CreatedAt "15:04"  }} </span>
            <div class="arrow right" id="a{{$item.ID}}"  onclick="openPost({{$item.ID}})"> ▼</div>

            <span class="head">
                        <span class="site_ico">
                                {{ if $item.Site.Ico }}
                                    <a href="/site/{{$item.Site.ID}}">
                                <img src="{{$item.Site.Ico}}"></a>
                                {{ end }}

                        </span>
                <span class="title"><a target="_blank"
                                       href="{{urlfor "MainController.Link"}}/{{$item.ID}}"
                                       >{{$item.Title}}</a></span>
            </span>

            <p style="display: none;" id="p{{$item.ID}}">{{$item.Description}}</p>
        </div>
    {{ end }}
    <style>
    .item {
      position: relative;
      display: block;
    }
    .item:hover:after {
     content: attr(data-title);
         position: absolute; /* Абсолютное позиционирование */
        left: 10%; top: -50%; /* Положение подсказки */
        z-index: 1; /* Отображаем подсказку поверх других элементов */
        background: rgba(230,255,255,0.9); /* Полупрозрачный цвет фона */
        font-family: Arial, sans-serif; /* Гарнитура шрифта */
        font-size: 12px; /* Размер текста подсказки */
        padding: 5px 15px; /* Поля */
        border: 1px solid #333; /* Параметры рамки */
    }
    </style>
</div>
