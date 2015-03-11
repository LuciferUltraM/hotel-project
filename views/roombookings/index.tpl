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
            <th>Grand Total</th>
            <th>Status</th>
        </thead>
        {{< range $index, $rb := .RoomBookings >}}
        <tr>
          <td><a href="/roombooking/{{< $rb.RoomBookingNo >}}">{{< $rb.RoomBookingNo >}}</a></td>
          <td>[{{< range $rm := $rb.Rooms >}}
            {{< $rm.RoomNo >}}
          {{< end >}}
          ]</td>
          <td>{{< $rb.CheckInDate.Format "2 January 2006" >}}</td>
          <td>{{< $rb.CheckOutDate.Format "2 January 2006" >}}</td>
          <td>{{< $rb.NightAmount >}}</td>
          <td>{{< $rb.GrandTotal >}}</td>
          <td>{{< $rb.Status >}}</td>
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
