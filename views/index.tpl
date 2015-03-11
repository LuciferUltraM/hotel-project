{{< template "layout.tpl" . >}}

{{< define "title" >}}
Home
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div class="panel panel-primary">
    <div class="panel-heading">
        <h3 class="panel-title">Search</h3>
    </div>
    <div class="panel-body">
    <form class="form-horizontal">
    <div class="form-group">
      <label for="inputEmail3" class="col-sm-2 control-label">Check-in</label>
      <div class="col-sm-10">
        <input ui-date="dateOptions" name="CheckInDate">
      </div>
    </div>
    <div class="form-group">
      <label for="inputEmail3" class="col-sm-2 control-label">Check-out</label>
      <div class="col-sm-10">
        <input ui-date="dateOptions" name="CheckOutDate">
      </div>
    </div>
    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" class="btn btn-default">Search</button>
      </div>
    </div>
    </form>
    </div>
  </div>

  <div>
    <h4>Results</h4>
    <table class="table table-striped">
        <thead>
            <th>Room No.</th>
            <th>Floor</th>
            <th>Room Type</th>
            <th>Rate per room per night</th>
            <th></th>
        </thead>
        {{< range .Rooms >}}
        <tr>
          <td>{{< .RoomNo >}}</td>
          <td>{{< .Floor >}}</td>
          <td>{{< .RoomType.Name >}}</td>
          <td>{{< .RoomType.Rate >}}</td>
          <td><input type="button" class="btn btn-primary" value="Book" /></td>
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
