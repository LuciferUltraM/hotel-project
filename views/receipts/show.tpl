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
        {{< .RoomBooking.Status >}}
      </div>
    </div>
    </div>
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
