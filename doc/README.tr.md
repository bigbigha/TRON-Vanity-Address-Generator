# TRON Vanity Address Generator

Hafif, **açık kaynak** bir TRON vanity (özel desenli) adres üreteci.  
**Güvenli**, **tamamen çevrimdışı** çalışan ve özel ön ek / son ek içeren TRON adresleri üretmek için tasarlandı.  
Basit, hızlı ve geliştirici dostu.

## Diller

- [English](../README.md))
- [Русский](README.ru.md)
- [简体中文](README.zh-CN.md)
- [Türkçe](docs/README.tr.md)
- [ไทย](README.th.md)
- [Español](README.es.md)

## Özellikler

- **Tamamen çevrimdışı**: Ağ isteği yok, telemetri yok, otomatik yükleme yok.
- **Açık kaynak**: Küçük ve odaklı bir kod tabanı, denetlemesi kolay.
- **Sadece CPU**: Tek statik ikili dosya ile Linux, macOS ve Windows üzerinde çalışır.
- **Çok iş parçacıklı**: Tüm kullanılabilir CPU çekirdeklerini kullanarak daha hızlı üretim.
- **Güvenli rastgelelik**: Kriptografik olarak güvenli rastgele sayı üretimi kullanır.

## İndir

Go veya Git kurmanıza gerek yok.  
Hazır derlenmiş ikili dosyaları GitHub Releases sayfasından indirebilirsiniz:

- En son sürüm:  
  https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases/latest

Releases sayfasında, işletim sisteminize uygun dosyayı indirin:

- **Windows**: `tronvanity-windows-amd64.exe`
- **macOS (Intel)**: `tronvanity-darwin-amd64`
- **macOS (Apple Silicon)**: `tronvanity-darwin-arm64`
- **Linux**: `tronvanity-linux-amd64`

İndirilen dosyaların değiştirilmediğinden emin olmak için `SHA256SUMS.txt` dosyasını da indirip kontrol etmeniz önerilir.

## Yardımcı Betikler (Komut satırı kullanmak istemeyenler için)

Eğer doğrudan komut satırı kullanmak istemiyorsanız, depo içindeki yardımcı betikleri kullanabilirsiniz.

Ayrıntılı açıklama: `scripts/README.tr.md`

### Windows Yardımcı Betiği (kısa özet)

1. Releases sayfasından `tronvanity-windows-amd64.exe` dosyasını indirin.
2. `scripts` klasöründen `windows-run.bat` dosyasını indirin.
3. Her ikisini de **aynı klasöre** koyun.
4. `windows-run.bat` dosyasına çift tıklayın.
5. Ekrandaki talimatları izleyerek ön ek / son ek ve üretilecek adres sayısını girin.

### macOS / Linux Yardımcı Betiği (kısa özet)

1. Sisteminiz için uygun ikili dosyayı indirin:
    - `tronvanity-darwin-amd64` (Intel Mac)
    - `tronvanity-darwin-arm64` (Apple Silicon Mac)
    - `tronvanity-linux-amd64` (Linux)
2. `scripts` klasöründen `unix-run.sh` dosyasını indirin.
3. İkili dosya ile `unix-run.sh` dosyasını **aynı klasöre** koyun.
4. Terminal’de bu klasöre geçin, `chmod +x unix-run.sh` çalıştırın.
5. `./unix-run.sh` komutuyla betiği başlatın ve ekrandaki talimatları izleyin.

Bu betikler yalnızca yerel ikili dosyayı çağıran basit sarmalayıcılardır; internete bağlanmaz, veri göndermez veya sistem dosyalarını değiştirmez.

## Geliştiriciler için Kurulum

```bash
# Depoyu klonlayın
git clone https://github.com/bigbigha/TRON-Vanity-Address-Generator.git
cd TRON-Vanity-Address-Generator

# İkili dosyayı derleyin
go build -o tronvanity ./cmd/tronvanity

# Örnek çalışma
./tronvanity --prefix T888
```

## Teknik olmayan kullanıcılar için Hızlı Başlangıç

### Windows

1. Releases sayfasından `tronvanity-windows-amd64.exe` dosyasını indirin.
2. Dosyayı istediğiniz bir klasöre koyun (veya yukarıdaki `windows-run.bat` ile birlikte aynı klasöre yerleştirin).
3. Komut İstemi’ni (Command Prompt) açın ve bu klasöre geçin.
4. Örnek komutlar:

```cmd
:: T888 ile başlayan adres üret
tronvanity-windows-amd64.exe --prefix T888

:: 9999 ile biten adres üret
tronvanity-windows-amd64.exe --suffix 9999
```

İsterseniz aynı klasördeki `windows-run.bat` yardımcı betiğini kullanarak, komutları yazmadan da etkileşimli bir şekilde kullanabilirsiniz.

### macOS

1. Uygun ikili dosyayı indirin:
    - Intel Mac: `tronvanity-darwin-amd64`
    - Apple Silicon: `tronvanity-darwin-arm64`
2. Terminal’i açın ve dosyanın bulunduğu klasöre geçin.
3. İlk sefer için çalıştırma izni verin:

```bash
chmod +x tronvanity-darwin-amd64      # Intel
chmod +x tronvanity-darwin-arm64      # Apple Silicon
```

4. Örnek komutlar:

```bash
# T888 ile başlayan adres
./tronvanity-darwin-amd64 --prefix T888

# 9999 ile biten adres
./tronvanity-darwin-arm64 --suffix 9999
```

İsterseniz `unix-run.sh` betiğini kullanarak daha basit, etkileşimli bir arayüzle çalışabilirsiniz.

### Linux

1. `tronvanity-linux-amd64` dosyasını indirin.
2. Terminal’de dosyanın bulunduğu klasöre geçin.
3. Çalıştırma izni verin:

```bash
chmod +x tronvanity-linux-amd64
```

4. Örnek komutlar:

```bash
# T888 ile başlayan adres
./tronvanity-linux-amd64 --prefix T888

# 9999 ile biten adres
./tronvanity-linux-amd64 --suffix 9999
```

> Not: Araç tamamen çevrimdışı çalışır. Daha yüksek güvenlik için, ikili dosyayı internete bağlı olmayan (air-gapped) bir makineye kopyalayıp orada kullanabilirsiniz.

## Kullanım

Temel komut satırı örnekleri:

```bash
# Sadece ön ek eşleştir
tronvanity --prefix T888

# Sadece son ek eşleştir
tronvanity --suffix 9999

# Hem ön ek hem son ek, 8 worker ile çalış, 3 eşleşme bulunca çık
tronvanity --prefix T666 --suffix 888 --workers 8 --count 3

# T123 ile başlayan 5 adres üret
tronvanity --prefix T123 --count 5
```

### Komut satırı seçenekleri

- `--prefix`: Eşleştirilecek ön ek (ör. `T888`).
- `--suffix`: Eşleştirilecek son ek (ör. `9999`).
- `--workers`: Worker goroutine sayısı (varsayılan: CPU çekirdek sayısı).
- `--count`: Çıkmadan önce bulunacak eşleşme sayısı (varsayılan: 1).

## Güvenlik İpuçları

1. **Çevrimdışı üretim**: Yüksek değerli adresler için aracı mümkün olduğunca internete kapalı bir makinede çalıştırın.
2. **Adresi kontrol edin**: Fon göndermeden önce, üretilen adresin istediğiniz ön ek / son eki içerdiğinden emin olun.
3. **Özel anahtarları koruyun**: Özel anahtarları donanım cüzdanında, parola yöneticisinde veya çevrimdışı şifreli yedekte saklayın. Güvenmediğiniz web sitelerine veya uygulamalara asla özel anahtar yapıştırmayın.
4. **Önce küçük tutarlarla test edin**: Büyük tutarlar göndermeden önce, küçük tutarlarla alım / gönderim testleri yapın.

## Nasıl Çalışır? (Kısa Özet)

Bu araç, TRON vanity adresleri bulmak için yerel bir brute-force araması yapar:

1. Go’nun kriptografik olarak güvenli rastgele sayı üreticisi ile 32 baytlık özel anahtar üretir.
2. Aynı TRON ve Ethereum’un kullandığı secp256k1 eliptik eğrisi ile karşılık gelen sıkıştırılmamış (uncompressed) genel anahtarı hesaplar.
3. Genel anahtarın baştaki `0x04` baytını atar, kalan kısma Keccak-256 uygular ve son 20 baytı alır.
4. Başına TRON ana ağ ön eki `0x41` ekleyerek 21 baytlık adres verisini elde eder.
5. Bu veriyi Base58Check ile kodlayarak standart `T...` TRON adresine dönüştürür.
6. Birden fazla worker bu süreci paralel olarak tekrarlar ve her üretilen adresi, istediğiniz ön ek / son ek ile eşleşip eşleşmediğini kontrol eder. Eşleşme bulunduğunda, adres ve özel anahtar CLI üzerinden yazdırılır.

Tüm bu adımlar **yalnızca yerel makinenizde** çalışır; herhangi bir ağ isteği yapılmaz ve özel anahtarlar hiçbir yere gönderilmez.

## Lisans

MIT

## Sorumluluk Reddi

Bu yazılım “olduğu gibi” sunulmaktadır ve hiçbir garanti verilmez.  
Kullanım riski tamamen size aittir.  
Gerçek fonlarla kullanmadan önce üretilen adresleri ve özel anahtarları dikkatlice doğrulayın ve güvenli depolama için gerekli önlemleri aldığınızdan emin olun.