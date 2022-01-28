<?php
$arr = [
    "config" => [
        "Name" => "/dev/serial/by-id/usb-FTDI_USB__-__Serial-if00-port0",
        "Baud" => 1200,
        "Parity" => 79,
        "Size" => 7,
        "StopBits" => 1
    ],
    "license" => [
        "licenseKey" => "f574283b71730179a3602d371df6aaaa94b7de2f234a5bf2c0b094e886f87f9c265825962328d5c35613a9d1c2438f91363a33067664c77208c7f782ce53ca2d0abc3890ce1f4422e913613a45f26f58c9d7594766b7c745af00d549bd0a42be7a3a1d6604a75ed52a2365f6ea59",
        "model" => 'MCM25',
        'licenseTerm' => '2022-01-22 00:00:00',
        'factoryNumber' => 'testttreaaassesssht'
    ]
];

echo json_encode($arr) . "\n";