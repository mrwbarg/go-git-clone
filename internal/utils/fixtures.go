package utils

var KVLMFixture []byte = []byte(`tree 29ff16c9c14e2652b22f8b78bb08a5a07930c147
parent 206941306e8a8af65b66eaaaea388a7ae24d49a0
author Mauricio Barg <mbarg@email.com> 1527025023 +0200
author Peter Parker <pparker@email.com> 1527025023 +0200
committer Mauricio Barg <mbarg@email.com> 1527025044 +0200
gpgsig -----BEGIN PGP SIGNATURE-----
 
 iQIzBAABCAAdFiEExwXquOM8bWb4Q2zVGxM2FxoLkGQFAlsEjZQACgkQGxM2FxoL
 kGQdcBAAqPP+ln4nGDd2gETXjvOpOxLzIMEw4A9gU6CzWzm+oB8mEIKyaH0UFIPh
 rNUZ1j7/ZGFNeBDtT55LPdPIQw4KKlcf6kC8MPWP3qSu3xHqx12C5zyai2duFZUU
 wqOt9iCFCscFQYqKs3xsHI+ncQb+PGjVZA8+jPw7nrPIkeSXQV2aZb1E68wa2YIL
 3eYgTUKz34cB6tAq9YwHnZpyPx8UJCZGkshpJmgtZ3mCbtQaO17LoihnqPn4UOMr
 V75R/7FjSuPLS8NaZF4wfi52btXMSxO/u7GuoJkzJscP3p4qtwe6Rl9dc1XC8P7k
 NIbGZ5Yg5cEPcfmhgXFOhQZkD0yxcJqBUcoFpnp2vu5XJl2E5I/quIyVxUXi6O6c
 /obspcvace4wy8uO0bdVhc4nJ+Rla4InVSJaUaBeiHTW8kReSFYyMmDCzLjGIu1q
 doU61OM3Zv1ptsLu3gUE6GU27iWYj2RWN3e3HE4Sbd89IFwLXNdSuM0ifDLZk7AQ
 WBhRhipCCgZhkj9g2NEk7jRVslti1NdN5zoQLaJNqSwO1MtxTmJ15Ksk3QP6kfLB
 Q52UWybBzpaP9HEd4XnR+HuQ4k2K0ns2KgNImsNvIyFwbpMUyUWLMPimaV1DWUXo
 5SBjDB/V/W2JBFR+XKHFJeFwYhj7DD/ocsGr4ZMx/lgc8rjIBkI=
 =lgTX
 -----END PGP SIGNATURE----

With great power, 
comes great responsibility.`)

var ParentFixture []byte = []byte(`tree 29ff16c9c14e2652b22f8b78bb08a5a07930c147
author Mauricio Barg <mbarg@email.com> 1527025023 +0200
committer Mauricio Barg <mbarg@email.com> 1527025044 +0200

With great power, 
comes great responsibility.`)

var ChildFixture []byte = []byte(`tree 29ff16c9c14e2652b22f8b78bb08a5a07930c147
parent cc8bf0100229330779d631d0c21662860337fc01
author Mauricio Barg <mbarg@email.com> 1527025023 +0200
committer Mauricio Barg <mbarg@email.com> 1527025044 +0200

And then he died. `)

var LogFixture = `commit 3d304ce573788d99e2a15382184b1435ccd6f102
Author: Mauricio Barg <mbarg@email.com>

	And then he died. 

commit cc8bf0100229330779d631d0c21662860337fc01
Author: Mauricio Barg <mbarg@email.com>

	With great power, 
	comes great responsibility.

`
