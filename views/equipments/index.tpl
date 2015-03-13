{{< template "layout.tpl" . >}}

{{< define "title" >}}
Room
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div>
    <h4>Equipments</h4>
    <table class="table table-striped">
        <thead>
            <th>Name</th>
            <th>Price</th>
        </thead>
        {{< range .Equipments >}}
        <tr>
          <td>{{< .Name >}}</td>
          <td>{{< .Price >}}</td>
        </tr>
        {{< end >}}
      </table>
  </div>
</div>
{{< end >}}

{{< define "js" >}}
<script>
app = angular.module('HotelApp', []);
app.controller('HotelCtrl', function ($scope) {
});
</script>
{{< end >}}
