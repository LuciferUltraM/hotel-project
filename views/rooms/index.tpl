{{< template "layout.tpl" . >}}

{{< define "title" >}}
Room
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div>
    <h4>Rooms</h4>
    <table class="table table-striped">
        <thead>
            <th>Room No.</th>
            <th>Floor</th>
            <th>Room Type</th>
            <th>Rate per room per night</th>
            <th>Status</th>
        </thead>
        {{< range .Rooms >}}
        <tr>
          <td>{{< .RoomNo >}}</td>
          <td>{{< .Floor >}}</td>
          <td>{{< .RoomType.Name >}}</td>
          <td>{{< .RoomType.Rate >}}</td>
          <td>{{< .Status >}}</td>
        </tr>
        {{< end >}}
      </table>
  </div>
</div>
{{< end >}}

{{< define "js" >}}
<script>
app = angular.module('HotelApp', ['ui.date']);
app.controller('HotelCtrl', function ($scope) {
  $scope.dateOptions = {
          changeYear: true,
          changeMonth: true,
          dateFormat: 'yy-mm-dd',
  };
});
</script>
{{< end >}}
