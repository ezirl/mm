<div class="article">
    <div class="title-page"><h1>Сообщения</h1></div>
{{/*    <h3><a href="/kosmo/election"><button>РЕДАКТИРОВАТЬ БЛОК ВЫБОРЫ</button></a></h3>*/}}

<form action="/kosmo/addnews" method="POST">
   <label> Название: <input type="text" name="title"></label>
   <label> Ссылка: <input type="text" name="url"></label>
   <label> Краткое описание: <textarea name="description" id="" cols="60" rows="5"></textarea></label>
   <label>Сайт: 
       <select name="site">
        {{ range $item := .Sites }}
            <option value="{{$item.ID}}">{{$item.Name}}</option>
        {{end}}
       </select>
   </label>

    <button>Добавить новость</button>
</form>
    <hr>
    <div class="adm-contacts">
        <a href="/kosmo/contacts/spam"><button
            style="background-color: darkred; margin-bottom: 10px;">Удалить предпологаемый спам</button></a>
    {{ range $item := .Contacts}}
        <div> Site: <a href="{{ $item.Link }}">{{ $item.Name }}</a> | Zone: {{$item.Country}} |
            Rss: <a href="{{$item.Rss}}" title="{{$item.Rss}}">rss</a> |
            <a href="{{ urlfor "MainController.DelContact"}}/{{$item.ID}}">DELETE </a>
        </div>
        <div> Email: {{$item.Email}}  |  Time: {{$item.CreatedAt}}</div>
        <div> <strong>Massage:</strong> {{$item.Comm}}</div>
        <div><a href="{{ urlfor "MainController.DelContact"}}/{{$item.ID}}">DELETE </a></div>
        <hr>
    {{end}}
    </div>
</div>

<div class="article" style="margin-top: 30px">
    <div class="title-page"><h1>Сайты</h1></div>

<a href="{{ urlfor "SiteController.AddSite" }}"><button>Добавить сайт</button></a>
    category: 1 - politics; 2 - economics; 3 - administration; 4 - world news;
    5 - science;
    <table cellspacing="10px" align="center" border="1px" cellpadding="20px">
    <tr>
        <th>Site</th>
        <th>Zone</th>
        <th>Rss</th>
        <th>Category</th>
        <th>Обновить</th>
        <th>Удалить?</th>
    </tr>
    {{ range $item := .Sites}}
        <tr>
            <td><a href="{{ $item.Link }}">{{ $item.Name }}</a></td>
            <td>{{$item.Country.Name}}</td>
            <td> <a href="{{ $item.Rss }}">Rss</a> </td>
            <td>{{$item.Category}}</td>
            <td> <a href="{{urlfor "SiteController.Put"}}/{{$item.ID}}">Edit </a></td>
            <td> <a href="{{urlfor "SiteController.Del"}}/{{$item.ID}}">DELETE </a></td>
        </tr>
    {{end}}
</table>
</div>
