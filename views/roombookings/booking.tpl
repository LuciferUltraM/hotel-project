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
          <span class="pull-right">Check-in Date {{< .RoomBooking.CheckInDate.Format "2 January" >}} to {{< .RoomBooking.CheckOutDate.Format "2 January" >}}</span></h3>
    </div>
    <div class="panel-body">
    <form class="form-horizontal" id="SearchRoom" method="POST" action="/roombooking/{{< .RoomBooking.RoomBookingNo >}}">
    <div class="form-group">
      <label for="inputFirstName" class="col-sm-2 control-label">First Name</label>
      <div class="col-sm-10">
        <input id="inputFirstName" type="text" name="Firstname" value="{{< .RoomBooking.Firstname >}}">
      </div>
    </div>
    <div class="form-group">
      <label for="inputLastname" class="col-sm-2 control-label">Last Name</label>
      <div class="col-sm-10">
        <input id="inputLastname" type="text" name="Lastname" value="{{< .RoomBooking.Lastname >}}">
      </div>
    </div>
    <div class="form-group">
      <label for="inputCardID" class="col-sm-2 control-label"> Card ID / Passport ID</label>
      <div class="col-sm-10">
        <input id="inputCardID" type="text" name="CardID" value="{{< .RoomBooking.CardID >}}">
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-10">
        <button type="submit" class="btn btn-primary" name="action">Confirm</button>
        <button type="submit" class="btn btn-primary" name="action">Action</button>
      </div>
    </div>

    <table class="table table-striped">
        <thead>
            <th>Room No.</th>
            <th>Floor</th>
            <th>Room Type</th>
            <th>Extra Bed</th>
            <th>Rate per room per night</th>
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
          </tr>
          {{< end >}}
        </tbody>
        <tfoot>
          <tr>
            <td colspan="4">Night</td>
            <td>{{< .RoomBooking.NightAmount >}}</td>
          </tr>
          <tr>
            <td colspan="4">Amount</td>
            <td>{{< .RoomBooking.Amount >}}</td>
          </tr>
          <tr>
            <td colspan="4">Vat</td>
            <td>{{< .RoomBooking.Vat >}}</td>
          </tr>
          <tr>
            <td colspan="4">Grand Total</td>
            <td>{{< .RoomBooking.GrandTotal >}}</td>
          </tr>
        </tfoot>
      </table>
    </form>
    </div>
  </div>
</div>
{{< end >}}

{{< define "js" >}}
{{< end >}}
