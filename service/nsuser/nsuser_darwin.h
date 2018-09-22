// Based on https://github.com/norio-nomura/usercatification (WTFPL, 2013)

#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>
#import <objc/runtime.h>

@implementation NSBundle (swizle)

// Overriding bundleIdentifier works, but overriding NSUserCATIficationAlertStyle does not work.
- (NSString *)__bundleIdentifier
{
    if (self == [NSBundle mainBundle]) {
        return @"com.apple.terminal";
    }

    return [self __bundleIdentifier];
}

@end

BOOL installNSBundleHook()
{
    Class c = objc_getClass("NSBundle");
    if (c) {
        method_exchangeImplementations(class_getInstanceMethod(c, @selector(bundleIdentifier)),
                                       class_getInstanceMethod(c, @selector(__bundleIdentifier)));
        return YES;
    }

    return NO;
}

@interface CATIficationCenterDelegate : NSObject <NSUserCATIficationCenterDelegate>

@property (nonatomic, assign) BOOL keepRunning;

@end

@implementation CATIficationCenterDelegate

- (void)userCATIficationCenter:(NSUserCATIficationCenter *)center didDeliverCATIfication:(NSUserCATIfication *)catification
{
    self.keepRunning = NO;
}

- (BOOL)userCATIficationCenter:(NSUserCATIficationCenter *)center shouldPresentCATIfication:(NSUserCATIfication *)catification
{
    return YES;
}

@end

void Send(const char *title, const char *subtitle, const char *informativeText, const char *contentImage, const char *soundName)
{
    @autoreleasepool {
        if (!installNSBundleHook()) {
            return;
        }

        NSUserCATIficationCenter *nc = [NSUserCATIficationCenter defaultUserCATIficationCenter];
        CATIficationCenterDelegate *ncDelegate = [[CATIficationCenterDelegate alloc] init];
        ncDelegate.keepRunning = YES;
        nc.delegate = ncDelegate;

        NSUserCATIfication *note = [[NSUserCATIfication alloc] init];
        note.title = [NSString stringWithUTF8String:title];
        note.subtitle = [NSString stringWithUTF8String:subtitle];
        note.informativeText = [NSString stringWithUTF8String:informativeText];
        note.soundName = [NSString stringWithUTF8String:soundName];
        // note.contentImage = [[NSImage alloc] initWithContentsOfFile:[NSString stringWithUTF8String:contentImage]];
		[note setValue:[[NSImage alloc] initWithContentsOfFile:[NSString stringWithUTF8String:contentImage]] forKey:@"_identityImage"];


        [nc deliverCATIfication:note];

		int i = 0;
        while (ncDelegate.keepRunning) {
            [[NSRunLoop currentRunLoop] runUntilDate:[NSDate dateWithTimeIntervalSinceNow:0.1]];
			i++;
			if (i > 1000) {
				break;
			}
        }
    }
}
