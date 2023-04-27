<div class="article">
    <div class="title-page"><h1>Добавить СМИ</h1></div>
    <form action="/contact" method="post">
        <input type="text" name="name" placeholder="Название СМИ">
        <input type="text" name="rss" placeholder="Ссылка на RSS канал">
        <select name="country" id="">
            {{ range $item := .Country }}
                <option value="{{$item.Name}}">{{$item.Name}}</option>
            {{end}}
        </select>
        <label for="email">
            <input type="text" name="email" placeholder="email@com.com">
        </label>
        <textarea name="comm" placeholder="комментарий" cols="50" rows="5"></textarea>
        <button class="gradient-link">Отправить</button>
    </form>

</div>