{{< template "layout.tpl" . >}}

{{< define "title" >}}
Room
{{< end >}}

{{< define "css" >}}

{{< end >}}

{{< define "content" >}}
<div ng-controller="HotelCtrl">
  <div>
    <h4>Receipt</h4>
    <table class="table table-striped">
        <thead>
            <th>Receipt No.</th>
            <th>RoomBooking No</th>
            <th>Customer Name</th>
            <th>Receipt Date</th>
            <th>Type</th>
            <th>Status</th>
            <th>Amount</th>
        </thead>
        {{< range .Receipts >}}
        <tr>
          <td><a href="/receipt/{{< .ReceiptNo >}}">{{< .ReceiptNo >}}</td>
          <td><a href="/roombooking/{{< .RoomBooking.RoomBookingNo >}}">{{< .RoomBooking.RoomBookingNo >}}</a></td>
          <td>{{< .RoomBooking.FirstName >}} {{< .RoomBooking.LastName >}}</td>
          <td>{{< .ReceiptDate.Format "2 January 2006 at 3:04pm" >}}</td>
          <td>{{< .Type >}}</td>
          <td>{{< .Status >}}</td>
          <td>{{< .Amount >}}</td>
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
