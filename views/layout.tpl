<!DOCTYPE html>

<html lang="ru">
<head>
    <title>Massmedia24 - новостной агрегатор</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="google-site-verification" content="e40isX931iRKigltNju92yyNsiV429mD6cJPB_nMess"/>
    <meta name="yandex-verification" content="91075674383ad076"/>
		<meta property="og:image" content="/static/img/asd.png" />
    <meta name="telderi" content="966e46cd7fe8881fb9fd6e23eaeab90a" />
    <link rel="stylesheet" href="/static/css/main.css">
    <link rel="stylesheet" href="/static/css/bootstrap-grid.min.css">
    <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/normalize/8.0.1/normalize.min.css"
          integrity="sha256-l85OmPOjvil/SOvVt3HnSSjzF1TUMyT9eV0c2BzEGzU=" crossorigin="anonymous"/>
{{/*    adsense google*/}}


    <!-- Yandex.Metrika counter -->
    <script type="text/javascript">
        (function (m, e, t, r, i, k, a) {
            m[i] = m[i] || function () {
                (m[i].a = m[i].a || []).push(arguments)
            };
            m[i].l = 1 * new Date();
            k = e.createElement(t), a = e.getElementsByTagName(t)[0], k.async = 1, k.src = r, a.parentNode.insertBefore(k, a)
        })
        (window, document, "script", "https://mc.yandex.ru/metrika/tag.js", "ym");

        ym(52245700, "init", {
            id: 52245700,
            clickmap: true,
            trackLinks: true,
            accurateTrackBounce: true,
            webvisor: true
        });
    </script>
    <noscript>
        <div><img src="https://mc.yandex.ru/watch/52245700" style="position:absolute; left:-9999px;" alt=""/></div>
    </noscript>
    <!-- /Yandex.Metrika counter -->
    <link rel="shortcut icon" href="/static/img/favicon.ico" type="image/png">
</head>

<body>

<header class="">
    <div class="zemlya">
    </div>
    <div class="container">
        <div class="header row">
            <div class="col-3">
                <span class="logo"><a href="/"><img src="/static/img/massmedia.png" width="100%" alt=""></a></span>
            </div>
            <div class="col-9">

               <ul class="menu head-menu right lang">
                    <li><a href=# class=" link-lang">
                            <img src="/static/img/koleso.png"
                                 class="ico-lang">
                            Выбрать регион</a>
                        <ul class="submenu">
                            <li><a href="/">Все</a></li>
                            {{ range $item := .countries }}
                                <li><a href="/{{$item.Link}}">{{$item.Name}}</a></li>
                            {{end}}
                        </ul>
                    </li>

                </ul>

                 <ul class="menu head-menu right lang">
                    <li>
                            {{ if .country}} {{.country.Name}} {{else}} Все {{end}}

                    </li>

                </ul>


            </div>

        </div>
    </div>
</header>


<div class="container content">
    <div class="row">
        <sidebar class="col-xl-3 sidebar">

            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Популярное за 24 часа</div>
                <div class="sidebar-item__section">
                    <div class="politics">
                        {{ if .popular }}
                        <div class="politic-items">

                            {{ range $item := .popular }}

                                <div class="popular-item">
                                    <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                        <div>{{$item.Title}}</div>
                                        <div>
                                            <span class="popular-item-site">
                                                {{if $item.Site.Name }} {{$item.Site.Name}} {{else}}
                                                    {{$item.SiteName}} {{end}}</span>
                                            <span class="popular-item-site" style="float: right; margin-top: 4px;">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                        </div>
                                    </a>
                                </div>
                            {{end}}
                        </div>
                        {{end}}
                    </div>
                </div>
            </div>

            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Оф источники<span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    {{ if .elections }}
                        <div class="politic-items">

                            {{ range $item := .elections }}

                                <div class="popular-item">
                                    <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                        <div>{{$item.Title}}</div>
                                        <div>
                                            <span class="popular-item-site">{{$item.Site.Name}}</span>
                                            <span class="popular-item-site" style="float: right; margin-top: 4px;">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                        </div>
                                    </a>
                                </div>
                            {{end}}
                        </div>
                    {{end}}
                </div>

            </div>
            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Политика <span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    {{ if .politic }}
                        <div class="politics">

                            <div class="politic-items">

                                {{ range $item := .politic }}

                                    <div class="popular-item">
                                        <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                            <div>{{$item.Title}}</div>
                                            <div>
                                                <span class="popular-item-site">{{$item.Site.Name}}</span>
                                                <span class="popular-item-site" style="float: right; margin-top: 4px;">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                            </div>
                                        </a>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                    {{end}}
                </div>
            </div>
            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Экономика <span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    {{ if .economic }}
                        <div class="politics">

                            <div class="politic-items">

                                {{ range $item := .economic }}

                                    <div class="popular-item">
                                        <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                            <div>{{$item.Title}}</div>
                                            <div>
                                                <span class="popular-item-site">{{$item.Site.Name}}</span>
                                                <span class="popular-item-site-date">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                            </div>
                                        </a>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                    {{end}}
                </div>
            </div>

            <div class="sidebar-item">
                <div class="sidebar-item__head blue">В мире <span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    {{ if .world }}
                        <div class="politics">

                            <div class="politic-items">

                                {{ range $item := .world }}

                                    <div class="popular-item">
                                        <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                            <div>{{$item.Title}}</div>
                                            <div>
                                                <span class="popular-item-site">{{$item.Site.Name}}</span>
                                                <span class="popular-item-site-date">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                            </div>
                                        </a>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                    {{end}}
                </div>
            </div>
            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Наука <span id="drop">▼ </span></div>
                <div class="sidebar-item__section none">
                    {{ if .science }}
                        <div class="politics">

                            <div class="politic-items">

                                {{ range $item := .science }}

                                    <div class="popular-item">
                                        <a href="{{urlfor "MainController.Link"}}/{{$item.ID}}">
                                            <div>{{$item.Title}}</div>
                                            <div>
                                                <span class="popular-item-site">{{$item.Site.Name}}</span>
                                                <span class="popular-item-site-date">
                                                {{dateformat $item.CreatedAt "15:04 02.01"}}</span>
                                            </div>
                                        </a>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                    {{end}}
                </div>
            </div>
        </sidebar>
        <div class="col-xl-6" style="padding-left: 0; padding-right: 0;">

            <main>

                {{ .LayoutContent}}

            </main>
            {{if .paginator }}
                <button onclick="load()" class="next-page">⇓РАНЕЕ⇓</button>
                {{if gt .paginator.PageNums 1}}
                    <ul class="pagination pagination-sm">
                        {{if .paginator.HasPrev}}
                            <li><a href="{{.paginator.PageLinkPrev}}" id="prev">&lt;</a></li>
                        {{else}}
                            <li class="disabled"><a>&lt;</a></li>
                        {{end}}
                        {{range $index, $page := .paginator.Pages}}
                            <li{{if $.paginator.IsActive .}} class="active"{{end}}>
                                <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
                            </li>
                        {{end}}
                        {{if .paginator.HasNext}}
                            <li><a href="{{.paginator.PageLinkNext}}" id="next">&gt;</a></li>
                        {{else}}
                            <li class="disabled"><a>&gt;</a></li>
                        {{end}}
                    </ul>
                {{end}}
            {{end}}
        </div>



        <sidebar class="col-xl-3 sidebar">
            <div class="sidebar-item">
                <div class="sidebar-item__head blue">Погода <span id="drop"> ▲</span></div>
                <div class="sidebar-item__section" style="background-color: #254b79;">
                    <!-- Gismeteo informer START -->
                    <link rel="stylesheet" type="text/css"
                          href="https://s1.gismeteo.ua/static/css/informer2/gs_informerClient.min.css">
                    <div id="gsInformerID-5sRKsJB8Djy00a" class="gsInformer" style="width:100%;height:279px">
                        <div class="gsIContent">
                            <div id="cityLink">
                                <a href="https://www.gismeteo.ua/weather-moscow-4368/" title="Погода в Москве"
                                   target="_blank">
                                    <img src="https://s1.gismeteo.ua/static/images/gisloader.svg" width="24" height="24"
                                         alt="Погода в Москве">
                                </a>
                            </div>
                            <div class="gsLinks">
                                <table>
                                    <tr>
                                        <td>
                                            <div class="leftCol">
                                                <a href="https://www.gismeteo.ua" target="_blank" title="Погода">
                                                    <img alt="Погода"
                                                         src="https://s1.gismeteo.ua/static/images/informer2/logo-mini2.png"
                                                         align="middle" border="0" height="16" width="11"/>
                                                    <img src="https://s1.gismeteo.ua/static/images/gismeteo.svg"
                                                         border="0" align="middle" style="left: 5px; top:1px">
                                                </a>
                                            </div>
                                            <div class="rightCol">
                                                <a href="https://www.gismeteo.ua/weather-moscow-4368/14-days/"
                                                   target="_blank" title="Погода в Москве на 2 недели">
                                                    <img src="https://s1.gismeteo.ua/static/images/informer2/forecast-2weeks.ru.svg"
                                                         border="0" align="middle" style="top:auto"
                                                         alt="Погода в Москве на 2 недели">
                                                </a>
                                            </div>
                                        </td>
                                    </tr>
                                </table>
                            </div>
                        </div>
                    </div>
                    <script src="https://www.gismeteo.ua/ajax/getInformer/?hash=5sRKsJB8Djy00a"
                            type="text/javascript"></script>
                    <!-- Gismeteo informer END -->
                </div>
            </div>
            <div class="sidebar-item" style="background-color: transparent;">
                <div class="sidebar-item__head blue">Курс валют <span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    <iframe style="width:100%;border:0;overflow:hidden;background-color:transparent;height:358px"
                            scrolling="no"
                            src="https://fortrader.org/informers/getInformer?st=6&cat=7&title=%D0%9A%D1%83%D1%80%D1%81%D1%8B%20%D0%B2%D0%B0%D0%BB%D1%8E%D1%82%20%D0%A6%D0%91%20%D0%A0%D0%A4&texts=%7B%22toolTitle%22%3A%22%D0%92%D0%B0%D0%BB%D1%8E%D1%82%D0%B0%22%2C%22todayCourse%22%3A%22RUB%22%7D&mult=1&showGetBtn=0&hideHeader=0&hideDate=1&w=0&codes=1&colors=titleTextColor%3Dfff%2CtitleBackgroundColor%3D254b79%2CthTextColor%3Dfff%2CthBackgroundColor%3D254b79%2CsymbolTextColor%3Dfff%2CtableTextColor%3Dfff%2CborderTdColor%3D396396%2CtableBorderColor%3D254b79%2CprofitTextColor%3Dfff%2CprofitBackgroundColor%3D254b79%2ClossTextColor%3Dfff%2ClossBackgroundColor%3D254b79%2CoddBackgroundTrColor%3D254b79%2CevenBackgroundTrColor%3D254b79%2CdataTextColor%3Dfff%2CdataBackgroundColor%3D254b79%2CinformerLinkTextColor%3Dfff%2CinformerLinkBackgroundColor%3D254b79&items=2%2C21%2C27%2C49%2C14%2C1&columns=todayCourse&toCur=11111"></iframe>
                </div>
            </div>

            <div class="sidebar-item" style="background-color: transparent;">
                <div class="sidebar-item__head blue">Гороскоп <span id="drop"> ▼</span></div>
                <div class="sidebar-item__section none">
                    <div class="goroskop">
                        <div class="row" style="margin: 0 auto;">
                            <div class="col-6" style="text-align: center;">
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/aries/today.html">
                                        <img src="/static/img/goroskop/aries.png" alt=""></a>
                                    <figcaption>Овен</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/germini/today.html">
                                        <img src="/static/img/goroskop/gemini.png" alt=""></a>
                                    <figcaption>Близнецы</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/lion/today.html">
                                        <img src="/static/img/goroskop/lion.png" alt=""></a>

                                    <figcaption>Лев</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/libra/today.html">
                                        <img src="/static/img/goroskop/libra.png" alt=""></a>
                                    <figcaption>Весы</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/sagittarius/today.html">
                                        <img src="/static/img/goroskop/sagittarius.png" alt=""></a>
                                    <figcaption>Стрелец</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/aquarius/today.html">
                                        <img src="/static/img/goroskop/aquarius.png" alt=""></a>
                                    <figcaption>Водолей</figcaption>
                                </figure>
                            </div>

                            <div class="col-6" style="text-align: center;">
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/taurus/today.html">
                                        <img src="/static/img/goroskop/taurus.png" alt=""></a>
                                    <figcaption>Телец</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/cancer/today.html">
                                        <img src="/static/img/goroskop/cancer.png" alt=""></a>
                                    <figcaption>Рак</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/virgo/today.html">
                                        <img src="/static/img/goroskop/virgo.png" alt=""></a>
                                    <figcaption>Дева</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/scorpio/today.html">
                                        <img src="/static/img/goroskop/scorpio.png" alt=""></a>
                                    <figcaption>Скорпион</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/capricorn/today.html">
                                        <img src="/static/img/goroskop/capricorn.png" alt=""></a>
                                    <figcaption>Козерог</figcaption>
                                </figure>
                                <figure>
                                    <a target="_blank"
                                       href="https://orakul.com/horoscope/astrologic/general/pisces/today.html">
                                        <img src="/static/img/goroskop/pisces.png" alt=""></a>
                                    <figcaption>Рыбы</figcaption>
                                </figure>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <!-- sidebar -->
        </sidebar>
    </div>
<div id="toTop"><img src="/static/img/rocket-off.png" alt="" width="70px" id="rocket"></div>
</div>

<footer class="footer">
    <span class="nav-menu">
                <a href="/sites">Источники</a>
                <a href="/addsite">Добавить ресурс</a>
               <a href="/about">О портале</a>
                <a href="/contact">Контакты</a>
            </span>
    <div class="author">
        Official website:
        <a href="/">MassMedia24</a>
    </div>
</footer>
<div class="backdrop"></div>

<script
        src="https://code.jquery.com/jquery-3.3.1.min.js"
        integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8="
        crossorigin="anonymous"></script>
<script src="/static/js/main.js"></script>
</body>
</html>
