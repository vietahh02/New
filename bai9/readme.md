# Chuẩn bị

1. Cài đặt [Xampp](https://www.youtube.com/watch?v=g7hdB6f7Hv8)(windows), [MAMP](https://www.youtube.com/watch?v=duBxCquAyg0)(MacOS) _(với Ubuntu xem [ở đây](https://www.youtube.com/watch?v=3e3WuanaJqs) hoặc [ở đây](https://www.youtube.com/watch?v=dzGVgqDj5nI))_. _Có thể chọn bất kỳ 1 webserver khác nhưng phải hỗ trợ PHP 5.6 trở lên._
2. Chọn 1 IDE hỗ trợ PHP : notepadd++, Netbeans, Dreamwaver, PHPStorm, PHPDesigner, Eclipse
3. Khởi động Xampp/Mamp và viết thử một file index.php trong thư mục webroot có nội dung.

        <?php
        echo php_info();
        ?>

để thử xem server đã hoạt động chưa.
_Thư mục webroot xampp thường là `C:/xampp/htdocs` của MAMP là `/Applications/MAMP/htdocs/`_

# Bài tập

Thực hiện các bài tập 1->7(trừ bài 3) của [trainee003](https://github.com/colombo-trainee/trainee003) bằng PHP (Không dùng bất kỳ đoạn javascript nào).
Về input của các bài tập có thể nhập cứng trong code hoặc [truyền dữ liệu qua form/url](https://www.youtube.com/watch?v=RwN17wW-DAM)

# Bonus

### Bài 1. Sử dụng PHP giải quyết bài tập sau:

- Tên file : `paginator.php`
- Đầu vào

 - Số a
 - Số b
 - Số c
 - Nút **Hiển thị**

- Đầu ra : 

 - Level 1: Hiển thị các số chia hết cho b nhỏ hơn a và lớn hơn 0 thành các trang, mỗi trang có c số thỏa mãn. Nút **Trang trước**, **Trang sau** để hiển thị c số trước/sau c số hiện tại.
         
        [Trang trước][Trang sau]

 - Level 2: Hiển thị liên kết đến các trang như sau
 
        [Trang trước] [1] ... [3][4][**5**][6][7] ... [31] [Trang sau]

        [Trang trước] [1][2][**3**][4][5] ... [31] [Trang sau]
        
        [Trang trước] [1] ... [27][28][29][**30**][31] [Trang sau]

### Bài 2. Sử dụng PHP giải quyết bài tập sau:

- Tên file : `string.php`
- Đầu vào

 - Đoạn văn bản b
 - Ký tự a
 - Nút **Phân tích**

- Đầu ra : Sau khi ấn nút **Phân tích**

 - Số lần xuất hiện ký tự a
 - Danh sách các từ chứa ký tự a
 - Đoạn văn bản b với các ký tự a được bôi đậm
 - Đoạn văn bản b với các từ có ký tự a được bôi đậm
 
