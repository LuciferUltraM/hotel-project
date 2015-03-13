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
        <h3 class="panel-title">Receipt No.{{< .Receipt.ReceiptNo >}}</h3>
    </div>
    <div class="panel-body">
    <div class="form-group">
      <label for="inputFirstName" class="col-sm-2 control-label">Room Booking No.</label>
      <div class="col-sm-10">
        <h5><a href="/roombooking/{{< .Receipt.RoomBooking.RoomBookingNo >}}" >{{< .Receipt.RoomBooking.RoomBookingNo >}}</a></h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputFirstName" class="col-sm-2 control-label">First Name</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.RoomBooking.FirstName >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputLastname" class="col-sm-2 control-label">Last Name</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.RoomBooking.LastName >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputCardID" class="col-sm-2 control-label"> Card ID / Passport ID</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.RoomBooking.CardID >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label for="inputCardID" class="col-sm-2 control-label"> Contact No.</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.RoomBooking.ContactNo >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Receipt Date</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.ReceiptDate.Format "2 January 2006 at 3:04pm" >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Type</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.Type >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Amount</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.Amount >}}</h5>
      </div>
    </div>
    <div class="form-group">
      <label  class="col-sm-2 control-label">Receipt Status</label>
      <div class="col-sm-10">
        <h5>{{< .Receipt.Status >}}</h5>
      </div>
    </div>


    {{< if .Receipt.RoomBooking.Rooms >}}

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
          {{< range $index, $room := .Receipt.RoomBooking.Rooms >}}
          <tr>
            <td>{{< $room.RoomNo >}}</td>
            <td>{{< $room.Floor >}}</td>
            <td>{{< $room.RoomType.Name >}}</td>
            <td>
              {{<if index $.Receipt.RoomBooking.ExtraBeds $index>}}
                Yes
              {{< else >}}
                No
              {{< end >}}
            </td>
            <td>{{< $room.RoomType.Rate >}}
              {{<if index $.Receipt.RoomBooking.ExtraBeds $index>}}
                + {{< $.Receipt.RoomBooking.ExtraBedRate >}}
              {{< end >}}
            </td>
            <td>
              {{< $.Receipt.RoomBooking.NightAmount >}}
            </td>
            <td class="text-right">{{
              {{<if index $.Receipt.RoomBooking.ExtraBeds $index>}}
                ({{<$room.RoomType.Rate>}} + {{<$.Receipt.RoomBooking.ExtraBedRate>}}) * {{<$.Receipt.RoomBooking.NightAmount>}}
              {{< else >}}
                {{<$room.RoomType.Rate>}} * {{<$.Receipt.RoomBooking.NightAmount>}}
              {{< end >}}
              }}
            </td>
          </tr>
          {{< end >}}
        </tbody>
        <tfoot>
          <tr>
            <td colspan="6"><strong>Amount</strong></td>
            <td class="text-right"><strong>{{< .Receipt.RoomBooking.Amount >}}</strong></td>
          </tr>
          <tr>
            <td colspan="6"><strong>Vat</strong></td>
            <td class="text-right"><strong>{{< .Receipt.RoomBooking.Vat >}}</strong></td>
          </tr>
          <tr>
            <td colspan="6"><strong>Grand Total</strong></td>
            <td class="text-right"><strong>{{< .Receipt.RoomBooking.GrandTotal >}}</strong></td>
          </tr>
        </tfoot>
      </table>

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
