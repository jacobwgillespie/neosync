'use client';
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
} from '@/components/ui/form';
import { Switch } from '@/components/ui/switch';
import { ReactElement } from 'react';
import { useFormContext } from 'react-hook-form';
interface Props {
  isDisabled?: boolean;
}

export default function CustomPhoneNumberTransformerForm(
  props: Props
): ReactElement {
  const fc = useFormContext();

  const { isDisabled } = props;

  return (
    <div className="flex flex-col w-full space-y-4 pt-4">
      <FormField
        name={`config.config.value.preserveLength`}
        control={fc.control}
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
            <div className="space-y-0.5">
              <FormLabel>Preserve Length</FormLabel>
              <FormDescription>
                Set the length of the output phone number to be the same as the
                input
              </FormDescription>
            </div>
            <FormControl>
              <Switch
                checked={field.value}
                onCheckedChange={field.onChange}
                disabled={
                  fc.watch('config.config.value.includeHyphens') || isDisabled
                }
              />
            </FormControl>
          </FormItem>
        )}
      />
      <FormField
        name={`config.config.value.includeHyphens`}
        control={fc.control}
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
            <div className="space-y-0.5">
              <FormLabel>Include Hyphens</FormLabel>
              <FormDescription>
                Include hyphens in the output phone number. Note: this only
                works with 10 digit phone numbers.
              </FormDescription>
            </div>
            <FormControl>
              <Switch
                checked={field.value}
                onCheckedChange={field.onChange}
                disabled={
                  fc.watch('config.config.value.preserveLength') || isDisabled
                }
              />
            </FormControl>
          </FormItem>
        )}
      />
      <FormField
        name={`config.config.value.e164Format`}
        control={fc.control}
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
            <div className="space-y-0.5">
              <FormLabel>Format in E164 Format</FormLabel>
              <FormDescription>
                Format the output phone number in the E164 Format. For ex.
                +1892393573894
              </FormDescription>
            </div>
            <FormControl>
              <Switch
                checked={field.value}
                onCheckedChange={field.onChange}
                disabled={
                  fc.watch('config.config.value.includeHyphens') || isDisabled
                }
              />
            </FormControl>
          </FormItem>
        )}
      />
    </div>
  );
}
