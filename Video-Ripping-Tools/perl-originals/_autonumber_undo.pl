#!/perl -w

use strict;

use File::Copy;
open(my $LOG,'<','autonumber.log');

foreach my $row (<$LOG>) {
    chomp $row;
    my ($f1,$f2) = ($1,$2) if $row =~ /(.+) -> (.+)/;
    next unless $f1;
    next unless -f $f2;
    print "$f2 -> $f1\n";
    move($f2,$f1);
    }

close $LOG;
