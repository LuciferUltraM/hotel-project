{{< template "layout.tpl" . >}}

{{< define "title" >}}
Room Bookings
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div class="panel panel-primary">
    <div class="panel-heading">
        <h3 class="panel-title">Booking No.{{< .RoomBooking.RoomBookingNo >}}
          <span class="pull-right">Check-in Date {{< .RoomBooking.CheckInDate.Format "2 January 2006">}} to {{< .RoomBooking.CheckOutDate.Format "2 January 2006">}}</span></h3>
    </div>
    <div class="panel-body">
    <div class="form-group">
      <label for="inputFirstName" class="col-sm-2 control-label">First Name</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.FirstName >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputLastname" class="col-sm-2 control-label">Last Name</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.LastName >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputCardID" class="col-sm-2 control-label">Card ID / Passport ID</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.CardID >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputCardID" class="col-sm-2 control-label">Contact No.</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.ContactNo >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Status</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.Status >}}</h5>
      </div>
    </div>

      <form class="form-horizontal" method="POST" action="/roombooking/{{< .RoomBooking.RoomBookingNo >}}/checkout">
      <table class="table table-striped">
          <thead>
              <th>Name</th>
              <th>Price</th>
              <th>Quantity</th>
              <th>Price</th>
          </thead>
          {{< range .Equipments >}}
          <tr ng-init="Quantity[{{< .ID >}}] = 0;Price[{{< .ID >}}] = {{< .Price >}}">
            <td>{{< .Name >}}</td>
            <td>{{< .Price >}}</td>
            <td><input type='number' min="0" name='Quantity[]' ng-model="Quantity[{{< .ID >}}]" ng-change="UpdateGrandTotal()">
              <input  type='hidden' name='EquipmentID[]' value='{{< .ID >}}'  >
            </td>
            <td>
              {{Quantity[{{< .ID >}}] * Price[{{< .ID >}}]}}
            </td>
          </tr>
          {{< end >}}
          <tfoot>
            <tr>
              <td colspan="3"><strong>Grand Total</strong></td>
              <td><strong>{{GrandTotal}}<input  type='hidden' name='GrandTotal' value='{{GrandTotal}}'  ></strong></td>
            </tr>
          </tfoot>
        </table>
          <div class="form-group">
            <div class="col-sm-12">
              <button  class="btn btn-danger" type="submit">Check-out</button>
            </div>
          </div>
        </form>

    </div>
  </div>
</div>
{{< end >}}

{{< define "js" >}}
<script>
app = angular.module('HotelApp', []);
app.controller('HotelCtrl', function ($scope) {
  $scope.GrandTotal = 0;
  $scope.UpdateGrandTotal = function() {
    $scope.GrandTotal = 0;
    console.log($scope.Quantity)
    angular.forEach($scope.Quantity, function(value, key) {
      $scope.GrandTotal += $scope.Quantity[key] * $scope.Price[key];
    });
  };
});
</script>
{{< end >}}
