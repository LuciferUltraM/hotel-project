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
        <h3 class="panel-title">Search Available Rooms</h3>
    </div>
    <div class="panel-body">
    <form class="form-horizontal" id="SearchRoom" method="GET" action="/">
    <div class="form-group">
      <label for="inputCheckInDate" class="col-sm-2 control-label">Check-in Date</label>
      <div class="col-sm-10">
        <input id="inputCheckInDate" ui-date="checkInDateOptions" name="CheckInDate" ng-model="CheckInDate"  ng-change="changeCheckInDate()" value="{{< .SearchRoom.CheckInDate >}}">
      </div>
    </div>
    <div class="form-group">
      <label for="inputCheckOutDate" class="col-sm-2 control-label">Check-out Date</label>
      <div class="col-sm-10">
        <input id="inputCheckOutDate" ui-date="checkOutDateOptions" name="CheckOutDate" ng-model="CheckOutDate" value="{{< .SearchRoom.CheckOutDate >}}">
      </div>
    </div>
    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" class="btn btn-primary">Search</button>
      </div>
    </div>
    </form>
    </div>
  </div>

  {{< if .Rooms >}}
  <div>
    <form id="SelectdRoom" method="POST" action="/" >
      <input type="hidden" name="CheckInDate" value="{{< .SearchRoom.CheckInDate >}}">
      <input type="hidden" name="CheckOutDate" value="{{< .SearchRoom.CheckOutDate >}}">
      {{< if .User.FirstName >}}
        <div class="pull-right"><input type="submit" class="btn btn-success" value="Book" /></div>
      {{< end >}}
    <table class="table table-striped">
        <thead>
            <th></th>
            <th>Room No.</th>
            <th>Extra Bed</th>
            <th>Floor</th>
            <th>Room Type</th>
            <th>Rate per room per night</th>
        </thead>
        {{< range .Rooms >}}
        <tr>
          <td><input type="checkbox" name="RoomNo[]" value="{{< .RoomNo >}}" ng-model="no{{< .RoomNo >}}"/></td>
          <td>{{< .RoomNo >}}</td>
          <td>
            <div ng-show="no{{< .RoomNo >}}">
              <input type="checkbox" ng-model="extra{{< .RoomNo >}}" />
              <input  type='hidden' ng-disabled="no{{< .RoomNo >}} != true" value='{{extra{{< .RoomNo >}}}}' name='ExtraBed[]'>
            </div>
          </td>
          <td>{{< .Floor >}}</td>
          <td>{{< .RoomType.Name >}}</td>
          <td>{{< .RoomType.Rate >}}</td>
        </tr>
        {{< end >}}
      </table>
    </form>
  </div>
  {{< end >}}
</div>

{{< end >}}

{{< define "js" >}}
<script>
app = angular.module('HotelApp', ['ui.date']);
app.controller('HotelCtrl', function ($scope) {
  var today = new Date()
  var tomorrow = new Date(today)
  tomorrow.setDate(today.getDate()+1);

  $scope.checkInDateOptions = {
          changeYear: true,
          changeMonth: true,
          dateFormat: 'yy-mm-dd',
          minDate: today,
  };
  $scope.checkOutDateOptions = {
          changeYear: true,
          changeMonth: true,
          dateFormat: 'yy-mm-dd',
          minDate: tomorrow,
  };
  $scope.CheckInDate = today;
  $scope.CheckOutDate = tomorrow;

  $scope.changeCheckInDate = function() {
    $scope.checkOutDateOptions.minDate.setDate($scope.CheckInDate.getDate() + 1);
  }
});
</script>
{{< end >}}
