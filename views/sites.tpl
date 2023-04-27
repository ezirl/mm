<div class="article">
    <div class="title-page"><h1>Зарегистрированые СМИ</h1></div>
    <ul class="list-site">
        {{ range $item := .sites }}
            <li>
                <a href="/site/{{$item.ID}}">
                <span class="site_ico">
                                <img src="{{$item.Ico}}" title="{{$item.Name}}">
                </span>
                    {{$item.Name}}</a>
            </li>
        {{end}}
    </ul>
</div>