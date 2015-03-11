{{< template "layout.tpl" . >}}

{{< define "title" >}}
Room Bookings
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div>
    <h4>Room Bookings</h4>
    <table class="table table-striped">
        <thead>
            <th>Booking No.</th>
            <th>Room</th>
            <th>Check-in Date</th>
            <th>Check-out Date</th>
            <th>Night</th>
            <th>Amount</th>
            <th>Vat</th>
            <th>Grand Total</th>
        </thead>
        {{< range $index, $rb := .RoomBookings >}}
        <tr>
          <td>{{< $rb.RoomBookingNo >}}</td>
          <td>[{{< range $rm := $rb.Rooms >}}
            {{< $rm.RoomNo >}}
          {{< end >}}
          ]</td>
          <td>{{< $rb.CheckInDate >}}</td>
          <td>{{< $rb.CheckOutDate >}}</td>
          <td>{{< $rb.NightAmount >}}</td>
          <td>{{< $rb.Amount >}}</td>
          <td>{{< $rb.Vat >}}</td>
          <td>{{< $rb.GrandTotal >}}</td>
        </tr>
        {{< end >}}
      </table>
  </div>
</div>
{{< end >}}

{{< define "js" >}}
{{< end >}}
