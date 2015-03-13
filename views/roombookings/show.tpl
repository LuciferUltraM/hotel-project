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

    {{<if .RoomBooking.CheckIn>}}
    <div class="form-group">
      <label  class="col-sm-2 control-label">Check-in</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.CheckIn.CheckInDate.Format "2 January 2006 at 3:04pm" >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Deposit</label>
      <div class="col-sm-10">
        <h5>{{< .RoomBooking.CheckIn.Deposit>}}</h5>
      </div>
    </div>
    {{<end>}}

    <table class="table table-striped">
        <thead>
            <th>Room No.</th>
            <th>Floor</th>
            <th>Room Type</th>
            <th>Extra Bed</th>
            <th>Rate per night</th>
            <th>Night</th>
            <th>Amount (THB)</th>
        </thead>
        <tbody>
          {{< range $index, $room := .RoomBooking.Rooms >}}
          <tr>
            <td>{{< $room.RoomNo >}}</td>
            <td>{{< $room.Floor >}}</td>
            <td>{{< $room.RoomType.Name >}}</td>
            <td>
              {{<if index $.RoomBooking.ExtraBeds $index>}}
                Yes
              {{< else >}}
                No
              {{< end >}}
            </td>
            <td>{{< $room.RoomType.Rate >}}
              {{<if index $.RoomBooking.ExtraBeds $index>}}
                + {{< $.RoomBooking.ExtraBedRate >}}
              {{< end >}}
            </td>
            <td>
              {{< $.RoomBooking.NightAmount >}}
            </td>
            <td class="text-right">{{
              {{<if index $.RoomBooking.ExtraBeds $index>}}
                ({{<$room.RoomType.Rate>}} + {{<$.RoomBooking.ExtraBedRate>}}) * {{<$.RoomBooking.NightAmount>}}
              {{< else >}}
                {{<$room.RoomType.Rate>}} * {{<$.RoomBooking.NightAmount>}}
              {{< end >}}
              }}
            </td>
          </tr>
          {{< end >}}
        </tbody>
        <tfoot>
          <tr>
            <td colspan="6"><strong>Amount</strong></td>
            <td class="text-right"><strong>{{< .RoomBooking.Amount >}}</strong></td>
          </tr>
          <tr>
            <td colspan="6"><strong>Vat</strong></td>
            <td class="text-right"><strong>{{< .RoomBooking.Vat >}}</strong></td>
          </tr>
          <tr>
            <td colspan="6"><strong>Grand Total</strong></td>
            <td class="text-right"><strong>{{< .RoomBooking.GrandTotal >}}</strong></td>
          </tr>
        </tfoot>
      </table>
      {{< if and .User .RoomBookingStatus.IsNew >}}
      <form class="form-horizontal" method="POST" action="/roombooking/{{< .RoomBooking.RoomBookingNo >}}/payment">
        <div class="form-group">
          <label  class="col-sm-2 control-label"> Payment Option</label>
          <div class="col-sm-10">
            <button  class="btn btn-primary" name="PaymentOption" type="submit" value="Cash">Cash</button>
            <button  class="btn btn-success" name="PaymentOption" type="submit" value="Credit Card">Credit Card</button>
          </div>
        </div>
      </form>
      {{< end >}}
      {{< if and .User .RoomBookingStatus.IsSuccess >}}
      <form class="form-horizontal" method="POST" action="/roombooking/{{< .RoomBooking.RoomBookingNo >}}/checkin">
        <div class="form-group">
          <div class="col-sm-12">
            <button  class="btn btn-primary" type="submit">Check-in</button>
          </div>
        </div>
      </form>
      {{< end >}}
      {{< if and .User .RoomBookingStatus.IsCheckIn >}}
      <form class="form-horizontal" method="GET" action="/roombooking/{{< .RoomBooking.RoomBookingNo >}}/checkout">
        <div class="form-group">
          <div class="col-sm-12">
            <button  class="btn btn-danger" type="submit">Check-out</button>
          </div>
        </div>
      </form>
      {{< end >}}
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
