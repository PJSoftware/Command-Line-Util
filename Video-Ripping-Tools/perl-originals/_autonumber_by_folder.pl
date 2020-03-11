#!/perl -w

use strict;

use Cwd;
use File::Copy;
use misc qw{ pad };

my $S;
my $E;
my $cwd = cwd;

open(my $LOG,'>','autonumber.log');

foreach my $sfolder (<S*>) {
    next unless -d $sfolder;
    if ($sfolder =~ /^s(\d+)$/i) {
        $S = $1;
        $E = 1;
        chdir($sfolder);
        foreach my $dfolder (<d*>) {
            next unless -d $dfolder;
            chdir($dfolder);
            foreach my $file (<*.txt>) {
                if ($file =~ /S${S}E(\d+).txt/) {
                    my $eTemp = $1;
                    $E = $eTemp if $eTemp >= $E;
                    }
                }
            foreach my $file (<*.m4v>) {
                my $new = $file;
                if ($new =~ /-(\d+)\.m4v/) {
                    my $num = $1;
                    while (length($num)<3) { $num = "0$num"; }
                    $new =~ s/-(\d+)\.m4v/-$num.m4v/;
                    move($file,$new) if $new ne $file;
                    }
                }

            foreach my $file (<*.m4v>) {
                my $new = "S$S"."E".pad($E,2,'0').".m4v";
                print "$sfolder/$dfolder/$file -> $new\n";
                print $LOG "$sfolder/$dfolder/$file -> $new\n";
                move($file,"$cwd/$new");
                $E++;
                }
            chdir('..');
            }
        chdir('..');
        }
    elsif ($sfolder =~ /^sx$/i) {
        chdir($sfolder);
        foreach my $dfolder (<s*e*>) {
            if ($dfolder =~ /^s(\d+)e(\d+)$/i && !rmdir($dfolder)) {
                ($S,$E) = ($1,$2);
                chdir($dfolder);
                foreach my $file (<*.m4v>) {
                    my $new = $file;
                    if ($new =~ /-(\d+)\.m4v/) {
                        my $num = $1;
                        while (length($num)<3) { $num = "0$num"; }
                        $new =~ s/-(\d+)\.m4v/-$num.m4v/;
                        move($file,$new) if $new ne $file;
                        }
                    }

                foreach my $file (<*.m4v>) {
                    my $new = "S$S"."E".pad($E,2,'0').".m4v";
                    print "$sfolder/$dfolder/$file -> $new\n";
                    print $LOG "$sfolder/$dfolder/$file -> $new\n";
                    move($file,"$cwd/$new");
                    $E++;
                    }
                chdir('..');
                }
            }
        chdir('..');
        }
    }

close $LOG;
