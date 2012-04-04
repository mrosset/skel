# This is an example PKGBUILD file. Use this as a start to creating your own,
# and remove these comments. For more information, see 'man PKGBUILD'.
# NOTE: Please fill out the license field for your package! If it is unknown,
# then please put 'unknown'.

# Maintainer: Your Name <youremail@domain.com>
pkgname=skel
pkgver=0.1
pkgrel=1
epoch=
pkgdesc="skel is a prototype go package"
arch=(i686 x86_64)
license=('new BSD')
depends=(go)
source=("https://github.com/downloads/str1ngs/$pkgname/$pkgname-$pkgver.tar.gz")
md5sums=('bb40f313b25c3f40734a881940db4e95')

build() {
  unset GOPATH GOROOT
  cd "${srcdir}/${pkgname}-${pkgver}"
  go build
}

check() {
  cd "${srcdir}/${pkgname}-${pkgver}"
  go test
}

package() {
  unset GOPATH GOROOT
  GOPATH="${pkgdir}" go install
}

# vim:set ts=2 sw=2 et:
