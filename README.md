# Golang_MailGun
# Nếu Chỉ Copy Source code
1. Tạo file mail.go với nội dung:

<!-- 

package main
import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var yourDomain string = "your-domain-name" 
var privateAPIKey string = "your-private-key"

func main() {
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "recipient@example.com"
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
} 

-->

2. Chạy lệnh "go mod init myapp" -> "go mod tidy" để bổ sung thư viện cần thiết

3. Tạo tài khoản mailgun trong ở trang https://app.mailgun.com nếu chưa có tài khoản

4. Lấy your-domain-name ở link: https://app.mailgun.com/app/domains

5. Lấy your-private-key ở link: https://app.mailgun.com/app/account/security/api_keys

6. Authorized Recipients cho tài khoản mail người nhận và vào mail người nhận để xác thực cấp quyền cho mailgun

7. Nhập đủ các trường:
    sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "recipient@example.com"

8. Chạy "go run main.go" và vào mail kiểm tra

! Định dạng mail sender có thể là sender := "sender@example.com" hoặc "Mailgun Sandbox<your-domain-name>"

# Nếu Clone code
Bắt đầu từ bước 3
