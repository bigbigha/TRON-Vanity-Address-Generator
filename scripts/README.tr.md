# TRON Vanity Address Generator Yardımcı Betikleri

Bu betikler, teknik bilgisi olmayan kullanıcılar için, komut satırını doğrudan kullanmadan TRON vanity adresleri üretmek için basit bir arayüz sağlar.

## Windows Betiği

### Kullanım

1. [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) sayfasından `tronvanity-windows-amd64.exe` dosyasını indirin.
2. Bu klasörden `windows-run.bat` dosyasını indirin.
3. Her iki dosyayı da **aynı klasöre** yerleştirin.
4. `windows-run.bat` dosyasına çift tıklayarak çalıştırın.
5. Ekrandaki talimatları izleyerek istediğiniz ön ek (prefix), son ek (suffix) ve üretilecek adres sayısını girin.

### Özellikler

- Basit, metin tabanlı arayüz.
- Komut satırı bilgisi gerektirmez.
- Siz `Ctrl + C` ile çıkış yapana kadar otomatik olarak tekrar eder.

## macOS / Linux Betiği

### Kullanım

1. [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) sayfasından sisteminize uygun ikili dosyayı indirin:
    - `tronvanity-darwin-amd64` (Intel Mac)
    - `tronvanity-darwin-arm64` (Apple Silicon Mac)
    - `tronvanity-linux-amd64` (Linux)
2. Bu klasörden `unix-run.sh` dosyasını indirin.
3. İkili dosya ile `unix-run.sh` dosyasını **aynı klasöre** yerleştirin.
4. Terminal’i açın ve bu klasöre geçin.
5. Çalıştırma izni verin: `chmod +x unix-run.sh`
6. Betiği çalıştırın: `./unix-run.sh`
7. Ekrandaki talimatları izleyin.

### Özellikler

- İşletim sistemi ve mimarinizi otomatik olarak algılar.
- Basit, metin tabanlı arayüz.
- İleri seviye komut satırı bilgisi gerektirmez.
- İkili dosyaya çalıştırma iznini otomatik verir.
- Siz `Ctrl + C` ile çıkış yapana kadar otomatik olarak tekrar eder.

## Güvenlik Notu

Bu betikler, yalnızca yerel ikili dosyayı çağıran basit sarmalayıcılardır. Şunları **yapmazlar**:

- İnternete erişmezler.
- Herhangi bir veriyi yüklemezler veya göndermezler.
- Yazılım yüklemezler.
- Sistem dosyalarını değiştirmezler.

Hem betikler hem de ana ikili dosya, tamamen **çevrimdışı** olarak yalnızca yerel makinenizde çalışır.