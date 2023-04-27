<div class="article">
    <div class="title-page"><h1>Выборы</h1></div>

    <a href="{{ urlfor "AdminController.AddPostGet" }}">
        <button>Добавить запись</button>
    </a>
    <table cellspacing="10px" align="center" border="1px" cellpadding="20px" width="100%">
        <tr>
            <th>Title</th>
            <th>Site</th>
            <th>Удалить?</th>
        </tr>
        {{ range $item := .Elections}}
            <tr>
                <td>{{$item.ID}} | <a href="{{ $item.Link }}">{{ $item.Title }}</a></td>
                <th>{{$item.SiteName}}</th>
                <td><a href="{{urlfor "AdminController.DelPost"}}/{{$item.ID}}">DELETE</a></td>
            </tr>
        {{end}}
    </table>
</div>